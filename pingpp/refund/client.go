// Package refund provides the /refunds APIs
package refund

import (
	"encoding/json"
	"fmt"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"log"
	"net/url"
	"strconv"
)

//client is used to invoke /refunds APIs
type Client struct {
	B   pingpp.Backend
	Key string
}

//New refunds a charge previously created .
func New(ch string, params *pingpp.RefundParams) (*pingpp.Refund, error) {
	return getC().New(ch, params)
}

func (c Client) New(ch string, params *pingpp.RefundParams) (*pingpp.Refund, error) {
	body := &url.Values{
		"amount": {strconv.FormatUint(params.Amount, 10)},
	}
	paramsString, errs := json.Marshal(params)

	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("RefundParams Marshall Errors is : %q\n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of refund request to pingpp is :\n %v\n ", string(paramsString))
	}
	refund := &pingpp.Refund{}
	err := c.B.Call("POST", fmt.Sprintf("/charges/%v/refunds", ch), c.Key, body, paramsString, refund)
	return refund, err
}

// Get returns the details of a refund.
func Get(chid string, reid string) (*pingpp.Refund, error) {
	return getC().Get(chid, reid)
}

func (c Client) Get(chid string, reid string) (*pingpp.Refund, error) {
	var body *url.Values
	body = &url.Values{}
	refund := &pingpp.Refund{}
	err := c.B.Call("GET", fmt.Sprintf("/charges/%v/refunds/%v", chid, reid), c.Key, body, nil, refund)
	return refund, err
}

// List returns a list of refunds.

func List(chid string, params *pingpp.RefundListParams) *Iter {
	return getC().List(chid, params)
}

func (c Client) List(chid string, params *pingpp.RefundListParams) *Iter {
	body := &url.Values{}
	var lp *pingpp.ListParams

	params.AppendTo(body)
	lp = &params.ListParams

	return &Iter{pingpp.GetIter(lp, body, func(b url.Values) ([]interface{}, pingpp.ListMeta, error) {
		list := &pingpp.RefundList{}
		err := c.B.Call("GET", fmt.Sprintf("/charges/%v/refunds", chid), c.Key, &b, nil, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Refunds.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*pingpp.Iter
}

// Refund returns the most recent Refund
// visited by a call to Next.
func (i *Iter) Refund() *pingpp.Refund {
	return i.Current().(*pingpp.Refund)
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

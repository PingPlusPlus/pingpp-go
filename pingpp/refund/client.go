// Package refund provides the /refunds APIs
package refund

import (
	"fmt"
	"log"
	"net/url"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func New(ch string, params *pingpp.RefundParams) (*pingpp.Refund, error) {
	return getC().New(ch, params)
}

func (c Client) New(ch string, params *pingpp.RefundParams) (*pingpp.Refund, error) {

	paramsString, errs := pingpp.JsonEncode(params)

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
	err := c.B.Call("POST", fmt.Sprintf("/charges/%v/refunds", ch), c.Key, nil, paramsString, refund)
	return refund, err
}

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

type Iter struct {
	*pingpp.Iter
}

func (i *Iter) Refund() *pingpp.Refund {
	return i.Current().(*pingpp.Refund)
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

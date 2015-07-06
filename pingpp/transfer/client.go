package transfer

import (
	"encoding/json"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"log"
	"net/url"
	"strconv"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func New(params *pingpp.TransferParams) (*pingpp.Transfer, error) {
	return getC().New(params)
}

func (c Client) New(params *pingpp.TransferParams) (*pingpp.Transfer, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("ChargeParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of redEnvelope request to pingpp is :\n %v\n ", string(paramsString))
	}
	transfer := &pingpp.Transfer{}
	err := c.B.Call("POST", "/transfers", c.Key, nil, paramsString, transfer)
	return transfer, err
}

// Get returns the details of a redenvelope.
func Get(id string) (*pingpp.Transfer, error) {
	return getC().Get(id)
}

func (c Client) Get(id string) (*pingpp.Transfer, error) {
	var body *url.Values
	body = &url.Values{}
	transfer := &pingpp.Transfer{}
	err := c.B.Call("GET", "/transfers/"+id, c.Key, body, nil, transfer)
	return transfer, err
}

// List returns a list of transfer.
func List(params *pingpp.TransferListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *pingpp.TransferListParams) *Iter {
	type transferList struct {
		pingpp.ListMeta
		Values []*pingpp.Transfer `json:"data"`
	}

	var body *url.Values
	var lp *pingpp.ListParams

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}
		params.AppendTo(body)
		lp = &params.ListParams
	}

	return &Iter{pingpp.GetIter(lp, body, func(b url.Values) ([]interface{}, pingpp.ListMeta, error) {
		list := &transferList{}
		err := c.B.Call("GET", "/transfers", c.Key, &b, nil, list)

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

func (i *Iter) Transfer() *pingpp.Transfer {
	return i.Current().(*pingpp.Transfer)
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

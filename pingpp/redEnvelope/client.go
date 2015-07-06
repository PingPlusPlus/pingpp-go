package redEnvelope

import (
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"log"
	"net/url"
	"strconv"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func New(params *pingpp.RedEnvelopeParams) (*pingpp.RedEnvelope, error) {
	return getC().New(params)
}

func (c Client) New(params *pingpp.RedEnvelopeParams) (*pingpp.RedEnvelope, error) {
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
	redEnvelope := &pingpp.RedEnvelope{}
	err := c.B.Call("POST", "/red_envelopes", c.Key, nil, paramsString, redEnvelope)
	return redEnvelope, err
}

func Get(id string) (*pingpp.RedEnvelope, error) {
	return getC().Get(id)
}

func (c Client) Get(id string) (*pingpp.RedEnvelope, error) {
	var body *url.Values
	body = &url.Values{}
	redEnvelope := &pingpp.RedEnvelope{}
	err := c.B.Call("GET", "/red_envelopes/"+id, c.Key, body, nil, redEnvelope)
	return redEnvelope, err
}

func List(params *pingpp.RedEnvelopeListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *pingpp.RedEnvelopeListParams) *Iter {
	type redEnvelopeList struct {
		pingpp.ListMeta
		Values []*pingpp.RedEnvelope `json:"data"`
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
		list := &redEnvelopeList{}
		err := c.B.Call("GET", "/red_envelopes", c.Key, &b, nil, list)

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

func (i *Iter) RedEnvelope() *pingpp.RedEnvelope {
	return i.Current().(*pingpp.RedEnvelope)
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

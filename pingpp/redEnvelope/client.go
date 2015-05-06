package redEnvelope

import (
	"encoding/json"
	"log"
	"net/url"
	"pingpp/pingpp"
	"strconv"
)

// Client is used to invoke /red_envelopes APIs.

type Client struct {
	B   pingpp.Backend
	Key string
}

// New POSTs new redenvelope.
func New(params *pingpp.RedEnvelopeParams) (*pingpp.RedEnvelope, error) {
	return getC().New(params)
}

func (c Client) New(params *pingpp.RedEnvelopeParams) (*pingpp.RedEnvelope, error) {
	body := &url.Values{
		"amount": {strconv.FormatUint(params.Amount, 10)},
	}
	paramsString, errs := json.Marshal(params)
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
	err := c.B.Call("POST", "/red_envelopes", c.Key, body, paramsString, redEnvelope)
	return redEnvelope, err
}

// Get returns the details of a redenvelope.
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

// List returns a list of redenvelope.
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

// Iter is an iterator for lists of redenvelope.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*pingpp.Iter
}

// redenvelope returns the most recent RedEnvelope
// visited by a call to Next.
func (i *Iter) RedEnvelope() *pingpp.RedEnvelope {
	return i.Current().(*pingpp.RedEnvelope)
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

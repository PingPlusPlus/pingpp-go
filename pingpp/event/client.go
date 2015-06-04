package event

import (
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"log"
	"net/url"
	"strconv"
)

// Client is used to invoke /charges APIs.

type Client struct {
	B   pingpp.Backend
	Key string
}

// Get returns the details of a charge.
func Get(id string) (*pingpp.Event, error) {
	return getC().Get(id)
}

func (c Client) Get(id string) (*pingpp.Event, error) {
	var body *url.Values
	body = &url.Values{}
	eve := &pingpp.Event{}
	err := c.B.Call("GET", "/events/"+id, c.Key, body, nil, eve)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get Event error: %v\n", err)
		}
	}
	return eve, err
}

// List returns a list of charges.
func List(params *pingpp.EventListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *pingpp.EventListParams) *Iter {
	type eventList struct {
		pingpp.ListMeta
		Values []*pingpp.Event `json:"data"`
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
		list := &eventList{}
		err := c.B.Call("GET", "/events", c.Key, &b, nil, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Charges.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*pingpp.Iter
}

// Charge returns the most recent Charge
// visited by a call to Next.
func (i *Iter) Event() *pingpp.Event {
	return i.Current().(*pingpp.Event)
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

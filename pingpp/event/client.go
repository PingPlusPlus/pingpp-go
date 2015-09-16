package event

import (
	"log"
	"net/url"
	pingpp "pingpp-go/pingpp"
	"strconv"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

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

type Iter struct {
	*pingpp.Iter
}

func (i *Iter) Event() *pingpp.Event {
	return i.Current().(*pingpp.Event)
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

package event

import (
	"log"
	"net/url"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
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

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

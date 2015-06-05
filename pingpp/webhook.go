package pingpp

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"log"
)

func ParseWebhooks(webhooks []byte) (*Event, error) {
	var event Event
	js, err := simplejson.NewJson(webhooks)

	if err != nil {
		return nil, err
	}

	_, ok := js.CheckGet("object")
	if ok {
		object, errObject := js.Get("object").String()
		if errObject != nil {
			log.Printf("cannot get the value of key object: %v\n", errObject)
		} else if object == "event" {
			errs := json.Unmarshal(webhooks, &event)
			if errs != nil {
				log.Printf("cannot unmarshal event object: %v\n", errs)
			}
		}
	}
	return &event, err
}

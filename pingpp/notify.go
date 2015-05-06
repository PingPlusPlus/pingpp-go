package pingpp

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"log"
)

func ParseNotify(notify []byte) (*Charge, *Refund, error) {
	var charge Charge
	var refund Refund
	js, err := simplejson.NewJson(notify)
	if err != nil {
		// panic("json format error")
		return nil, nil, err
	}

	_, ok := js.CheckGet("object")
	if ok {
		object, errObject := js.Get("object").String()
		if errObject != nil {
			log.Printf("cannot get the value of key object: %v\n", errObject)
		}
		if object == "charge" {
			errCh := json.Unmarshal(notify, &charge)
			if errCh != nil {
				log.Printf("cannot unmarshal to charge: %v\n", errCh)
				return nil, nil, errCh
			}
			return &charge, nil, nil

		} else if object == "fefund" {
			errRe := json.Unmarshal(notify, &refund)
			if errRe != nil {
				log.Printf("cannot unmarshal to charge: %v\n", errRe)
				return nil, nil, errRe
			}
			return nil, &refund, nil
		}

	} else {
		log.Println("this is not a pingpp notify,because key object not exists")
		return nil, nil, err
	}
	return nil, nil, err
}

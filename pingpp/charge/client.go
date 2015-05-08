package charge

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
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

// New POSTs new charges.
func New(params *pingpp.ChargeParams) (*pingpp.Charge, error) {
	return getC().New(params)
}

func (c Client) New(params *pingpp.ChargeParams) (*pingpp.Charge, error) {
	body := &url.Values{
		"amount": {strconv.FormatUint(params.Amount, 10)},
	}

	body.Add("order_no", params.Order_no)
	appstring, err := json.Marshal(params.App)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	body.Add("app", string(appstring))
	body.Add("channel", params.Channel)
	body.Add("currency", params.Currency)
	body.Add("subject", params.Subject)
	body.Add("body", params.Body)
	body.Add("client_ip", params.Client_ip)
	extrastring, err1 := json.Marshal(params.Extra)
	if err1 != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err1)
		}
	}
	body.Add("extra", string(extrastring))

	if len(params.Metadata) > 0 {
		metastring, _ := json.Marshal(params.Metadata)
		body.Add("metadata", string(metastring))
	}

	if params.Time_expire > 0 {
		body.Add("time_expire", strconv.FormatUint(params.Time_expire, 10))
	}

	if len(params.Description) > 0 {
		body.Add("description", string(params.Description))
	}

	paramsString, errs := json.Marshal(params)
	if pingpp.LogLevel > 2 {
		log.Printf("params of charge request to pingpp is :\n %v\n ", string(paramsString))
	}

	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("ChargeParams Marshall Errors is : %q\n", errs)
		}
	}

	charge := &pingpp.Charge{}
	errch := c.B.Call("POST", "/charges", c.Key, body, paramsString, charge)
	if errch != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", errch)
		}
		return nil, errch
	} else {
		chs, err := json.Marshal(charge)
		if err != nil {
			if pingpp.LogLevel > 0 {
				log.Printf("cannot marshal charge json: %v\n", err)
			}
		}
		js, err1 := simplejson.NewJson(chs)
		if err1 != nil {
			if pingpp.LogLevel > 0 {
				log.Printf("cannot unmarshal charge json: %v\n", err1)
			}
		}
		channel, err2 := js.Get("channel").String()
		if err2 != nil {
			if pingpp.LogLevel > 0 {
				log.Printf("cannot get channel from charge : %v\n", err2)
			}
		}
		if channel == "wx" {
			timestamp, _ := js.Get("credential").Get("wx").Get("timeStamp").Float64()
			s, err3 := js.Get("credential").Get("wx").Map()
			if err3 != nil {
				if pingpp.LogLevel > 0 {
					log.Printf("decode error: get int failed! %v\n", err3)
				}
				return nil, err3
			} else {
				timestamps := int(timestamp)
				s["timeStamp"] = strconv.Itoa(timestamps)
				chs, errchs := json.Marshal(js)
				if errchs != nil {
					if pingpp.LogLevel > 0 {
						log.Printf("cannot marshel charge: %v\n", errchs)
					}
				}
				errcharge := json.Unmarshal(chs, &charge)
				if errcharge != nil {
					if pingpp.LogLevel > 0 {
						log.Printf("cannot unmarshal charge from unmarshel by simplejson : %v\n", errcharge)
					}
				}
				return charge, errcharge
			}
		} else {
			return charge, errch
		}
	}
}

// Get returns the details of a charge.
func Get(id string) (*pingpp.Charge, error) {
	return getC().Get(id)
}

func (c Client) Get(id string) (*pingpp.Charge, error) {
	var body *url.Values
	body = &url.Values{}
	charge := &pingpp.Charge{}
	err := c.B.Call("GET", "/charges/"+id, c.Key, body, nil, charge)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get Charge error: %v\n", err)
		}
	}
	return charge, err
}

// List returns a list of charges.
func List(params *pingpp.ChargeListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *pingpp.ChargeListParams) *Iter {
	type chargeList struct {
		pingpp.ListMeta
		Values []*pingpp.Charge `json:"data"`
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
		list := &chargeList{}
		err := c.B.Call("GET", "/charges", c.Key, &b, nil, list)

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
func (i *Iter) Charge() *pingpp.Charge {
	return i.Current().(*pingpp.Charge)
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

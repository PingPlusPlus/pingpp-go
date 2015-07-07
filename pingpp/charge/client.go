package charge

import (
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"log"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

// 发送 charge 请求
func New(params *pingpp.ChargeParams) (*pingpp.Charge, error) {
	return getC().New(params)
}

func (c Client) New(params *pingpp.ChargeParams) (*pingpp.Charge, error) {
	start := time.Now()
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("ChargeParams Marshall Errors is : %q\n", errs)
		}
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of charge request to pingpp is :\n %v\n ", string(paramsString))
	}

	charge := &pingpp.Charge{}
	errch := c.B.Call("POST", "/charges", c.Key, nil, paramsString, charge)
	if errch != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", errch)
		}
		return nil, errch
	}
	if pingpp.LogLevel > 2 {
		log.Println("Charge completed in ", time.Since(start))
	}
	return charge, errch

}

//查询指定 charge 对象
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

// 查询 charge 列表
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

type Iter struct {
	*pingpp.Iter
}

func (i *Iter) Charge() *pingpp.Charge {
	return i.Current().(*pingpp.Charge)
}

package customer

import (
	"fmt"
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

// 发送 customer 请求
func New(params *pingpp.CustomerParams) (*pingpp.Customer, error) {
	return getC().New(params)
}

func (c Client) New(params *pingpp.CustomerParams) (*pingpp.Customer, error) {
	start := time.Now()
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("CustomerParams Marshall Errors is : %q\n", errs)
		}
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of card request to pingpp is :\n %v\n ", string(paramsString))
	}

	customer := &pingpp.Customer{}
	errch := c.B.Call("POST", "/customers", c.Key, nil, paramsString, customer)
	if errch != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", errch)
		}
		return nil, errch
	}
	if pingpp.LogLevel > 2 {
		log.Println("Card completed in ", time.Since(start))
	}
	return customer, errch

}

//查询指定 customer 对象
func Get(cus_id string) (*pingpp.Customer, error) {
	return getC().Get(cus_id)
}

func (c Client) Get(cus_id string) (*pingpp.Customer, error) {
	var body *url.Values
	body = &url.Values{}
	customer := &pingpp.Customer{}
	err := c.B.Call("GET", fmt.Sprintf("/customers/%v", cus_id), c.Key, body, nil, customer)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get Card error: %v\n", err)
		}
	}
	return customer, err
}

// 发送 customer 请求
func Update(cus_id string, params *pingpp.CustomerUpdateParams) (*pingpp.Customer, error) {
	return getC().Update(cus_id, params)
}

func (c Client) Update(cus_id string, params *pingpp.CustomerUpdateParams) (*pingpp.Customer, error) {
	start := time.Now()
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("CustomerParams Marshall Errors is : %q\n", errs)
		}
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of card request to pingpp is :\n %v\n ", string(paramsString))
	}

	customer := &pingpp.Customer{}
	errch := c.B.Call("PUT", fmt.Sprintf("/customers/%v", cus_id), c.Key, nil, paramsString, customer)
	if errch != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", errch)
		}
		return nil, errch
	}
	if pingpp.LogLevel > 2 {
		log.Println("Card completed in ", time.Since(start))
	}
	return customer, errch

}

//删除指定 customer 对象
func Delete(cus_id string) (map[string]interface{}, error) {
	return getC().Delete(cus_id)
}

func (c Client) Delete(cus_id string) (map[string]interface{}, error) {
	var body *url.Values
	body = &url.Values{}
	res := make(map[string]interface{})
	err := c.B.Call("DELETE", fmt.Sprintf("/customers/%v", cus_id), c.Key, body, nil, &res)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get Card error: %v\n", err)
		}
	}
	return res, err
}

// 查询 customer 列表
func List(params *pingpp.CustomerListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *pingpp.CustomerListParams) *Iter {
	type chargeList struct {
		pingpp.ListMeta
		Values []*pingpp.Customer `json:"data"`
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
		err := c.B.Call("GET", "/customers", c.Key, &b, nil, list)

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

func (i *Iter) Customer() *pingpp.Customer {
	return i.Current().(*pingpp.Customer)
}

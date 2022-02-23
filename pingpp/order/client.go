package order

import (
	"fmt"
	"log"
	"net/url"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func New(params *pingpp.OrderCreateParams) (*pingpp.Order, error) {
	return getC().New(params)
}

func (c Client) New(params *pingpp.OrderCreateParams) (*pingpp.Order, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("OrderCreateParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of create order is :\n %v\n ", string(paramsString))
	}

	order := &pingpp.Order{}
	err := c.B.Call("POST", "/orders", c.Key, nil, paramsString, order)
	return order, err
}

func Pay(id string, params *pingpp.OrderPayParams) (*pingpp.Order, error) {
	return getC().Pay(id, params)
}

func (c Client) Pay(id string, params *pingpp.OrderPayParams) (*pingpp.Order, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("OrderPayParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of order pay is :\n %v\n ", string(paramsString))
	}

	order := &pingpp.Order{}
	err := c.B.Call("POST", fmt.Sprintf("/orders/%s/pay", id), c.Key, nil, paramsString, order)
	return order, err
}

func Cancel(user, id string) (*pingpp.Order, error) {
	return getC().Cancel(user, id)
}
func (c Client) Cancel(user, id string) (*pingpp.Order, error) {
	params := struct {
		Status string `json:"status"`
		User   string `json:"user"`
	}{
		Status: "canceled",
		User:   user,
	}
	paramsString, _ := pingpp.JsonEncode(params)
	if pingpp.LogLevel > 2 {
		log.Printf("params of cancel order  is :\n %v\n ", string(paramsString))
	}

	order := &pingpp.Order{}
	err := c.B.Call("PUT", "/orders/"+id, c.Key, nil, paramsString, order)

	return order, err
}

func Get(id string) (*pingpp.Order, error) {
	return getC().Get(id)
}

func (c Client) Get(id string) (*pingpp.Order, error) {
	order := &pingpp.Order{}

	err := c.B.Call("GET", "/orders/"+id, c.Key, nil, nil, order)
	return order, err
}

func List(params *pingpp.PagingParams) (*pingpp.OrderList, error) {
	return getC().List(params)
}
func (c Client) List(params *pingpp.PagingParams) (*pingpp.OrderList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	orderList := &pingpp.OrderList{}
	err := c.B.Call("GET", "/orders", c.Key, body, nil, orderList)
	return orderList, err
}

func Charge(orderID, chargeID string) (*pingpp.Charge, error) {
	return getC().Charge(orderID, chargeID)
}

func (c Client) Charge(orderID, chargeID string) (*pingpp.Charge, error) {
	charge := &pingpp.Charge{}

	err := c.B.Call("GET", fmt.Sprintf("/orders/%s/charges/%s", orderID, chargeID), c.Key, nil, nil, charge)
	return charge, err
}

func ChargeList(orderID string, params *pingpp.PagingParams) (*pingpp.ChargeList, error) {
	return getC().ChargeList(orderID, params)
}

func (c Client) ChargeList(orderID string, params *pingpp.PagingParams) (*pingpp.ChargeList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	chargeList := &pingpp.ChargeList{}
	err := c.B.Call("GET", fmt.Sprintf("/orders/%s/charges", orderID), c.Key, body, nil, chargeList)
	return chargeList, err
}

func Update(id string, params *pingpp.OrderUpdateParams) (*pingpp.Order, error) {
	return getC().Update(id, params)
}

// Update 取消/更新 Order 对象
func (c Client) Update(id string, params *pingpp.OrderUpdateParams) (*pingpp.Order, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("OrderUpdateParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of update order is :\n %v\n ", string(paramsString))
	}

	order := &pingpp.Order{}
	err := c.B.Call("PUT", fmt.Sprintf("/orders/%s", id), c.Key, nil, paramsString, order)
	return order, err
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

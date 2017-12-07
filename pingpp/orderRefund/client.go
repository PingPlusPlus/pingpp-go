package orderRefund

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

func New(id string, params *pingpp.OrderRefundParams) (*pingpp.RefundList, error) {
	return getC().New(id, params)
}

func (c Client) New(id string, params *pingpp.OrderRefundParams) (*pingpp.RefundList, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("OrderRefundParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of orderRefund  is :\n %v\n ", string(paramsString))
	}

	orderRefund := &pingpp.RefundList{}
	err := c.B.Call("POST", fmt.Sprintf("/orders/%s/order_refunds", id), c.Key, nil, paramsString, orderRefund)
	return orderRefund, err
}

func Get(orderId, refundId string) (*pingpp.Refund, error) {
	return getC().Get(orderId, refundId)
}

func (c Client) Get(orderId, refundId string) (*pingpp.Refund, error) {
	refund := &pingpp.Refund{}

	err := c.B.Call("GET", fmt.Sprintf("/orders/%s/order_refunds/%s", orderId, refundId), c.Key, nil, nil, refund)
	return refund, err
}

func List(orderId string, params *pingpp.PagingParams) (*pingpp.RefundList, error) {
	return getC().List(orderId, params)
}

func (c Client) List(orderId string, params *pingpp.PagingParams) (*pingpp.RefundList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	orderRefundList := &pingpp.RefundList{}

	err := c.B.Call("GET", fmt.Sprintf("/orders/%s/order_refunds", orderId), c.Key, body, nil, orderRefundList)
	return orderRefundList, err
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

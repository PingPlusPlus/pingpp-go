package recharegRefund

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

func New(appID, id string, params *pingpp.RechargeRefundParams) (*pingpp.Refund, error) {
	return getC().New(appID, id, params)
}

func (c Client) New(appID, id string, params *pingpp.RechargeRefundParams) (*pingpp.Refund, error) {
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

	rechargeRefund := &pingpp.Refund{}
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/recharges/%s/refunds", appID, id), c.Key, nil, paramsString, rechargeRefund)
	return rechargeRefund, err
}

func Get(appID, rechargeID, refundID string) (*pingpp.Refund, error) {
	return getC().Get(appID, rechargeID, refundID)
}

func (c Client) Get(appID, rechargeID, refundID string) (*pingpp.Refund, error) {
	refund := &pingpp.Refund{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/recharges/%s/refunds/%s", appID, rechargeID, refundID), c.Key, nil, nil, refund)
	return refund, err
}

func List(appID, rechargeID string, params *pingpp.PagingParams) (*pingpp.RefundList, error) {
	return getC().List(appID, rechargeID, params)
}

func (c Client) List(appID, rechargeID string, params *pingpp.PagingParams) (*pingpp.RefundList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	rechargeRefundList := &pingpp.RefundList{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/recharges/%s/refunds", appID, rechargeID), c.Key, body, nil, rechargeRefundList)
	return rechargeRefundList, err
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

package pingpp

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

type RefundClient struct {
	key     string
	backend Backend
}

func getRefundClient(key string) *RefundClient {
	return &RefundClient{
		key:     key,
		backend: getBackend(),
	}
}

func (refundClient *RefundClient) setKey(key string) {
	refundClient.key = key
}

func (refundClient *RefundClient) new(params *RefundParams, charge_id string) (*Refund, error) {
	var refund Refund
	body := &url.Values{}
	if params.Amount > 0 && params.Amount <= 100 {
		body.Add("amount", strconv.FormatInt(params.Amount, 10))
	} else {
		body.Add("amount", strconv.FormatInt(params.Amount, 10))
	}
	if len(params.Description) > 0 {
		body.Add("description", params.Description)
	}
	resp_byte, err := refundClient.backend.Call("POST", charge_id+"/refunds", refundClient.key, body, &refund)
	json.Unmarshal(resp_byte, &refund)
	return &refund, err
}

func (refundClient *RefundClient) get(charge_id string, refund_id string) (*Refund, error) {
	var refund Refund
	urlvalues := &url.Values{}
	resp_byte, err := refundClient.backend.Call("GET", charge_id+"/refunds/"+refund_id, refundClient.key, urlvalues, &refund)
	json.Unmarshal(resp_byte, &refund)
	return &refund, err
}

func (refundClient *RefundClient) list(charge_id string, limit int64, start_after string, end_before string) (*RefundList, error) {
	var refundList RefundList
	body := &url.Values{}
	url := ""
	if charge_id == "" {
		err := errors.New("Lack of charge params:charge_id")
		return &refundList, err
	} else {
		url = url + charge_id + "/refunds?"
	}
	if limit > 0 && limit <= 100 {
		url = url + "limit=" + strconv.FormatInt(limit, 10) + "&"
	} else {
		url = url + "limit=10&"
	}
	if len(start_after) > 0 {
		url = url + "starting_after=" + start_after + "&"
	}
	if len(end_before) > 0 {
		url = url + "ending_before=" + end_before + "&"
	}
	resp_byte, err := refundClient.backend.Call("GET", url, refundClient.key, body, &refundList)
	json.Unmarshal(resp_byte, &refundList)
	return &refundList, err
}

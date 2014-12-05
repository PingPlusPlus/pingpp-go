package pingpp

import (
	"testing"
)

func TestChargeNew(t *testing.T) {

	chargeClient := GetChargeClient("sk_live_bPaPC8TyvHuD8uXjfPSKGa90")
	chargeParams := &ChargeParams{
		order_no:  "88888889",
		appid:     "app_9Kafv5qD0iP4jH48",
		channel:   "alipay",
		amount:    10,
		currency:  "cny",
		client_ip: "127.0.0.1",
		subject:   "test",
		body:      "bodysample",
	}

	charge, err := chargeClient.New(chargeParams)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(charge.Id))
	t.Log(string(charge.Object))
	t.Log(string(charge.App))
	t.Log(string(charge.Client_ip))
	t.Log(string(charge.Time_expire))
	t.Error(string(charge.Credential.AlipayCredential.AlipayOrderInfo))
}

func TestChargeList(t *testing.T) {
	var client *ChargeClient
	client = GetChargeClient("sk_live_bPaPC8TyvHuD8uXjfPSKGa90")
	params := &ChargeListParams{}
	target, err := client.List(params)
	if err != nil {
		t.Error(err)
	}
	t.Error(string(target.Url))
}

func TestChargeGet(t *testing.T) {
	client := GetChargeClient("sk_live_bPaPC8TyvHuD8uXjfPSKGa90")
	var id string
	id = "ch_n10ejHLibfzPPmf9i5u9uT4G"
	target, err := client.Get(id)
	if err != nil {
		t.Error(err)
	}
	t.Error(target)
}

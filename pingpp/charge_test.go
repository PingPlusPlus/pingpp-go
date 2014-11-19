package pingpp

import (
	"testing"
)

func TestChargeNew(t *testing.T) {

	chargeClient := GetChargeClient("sk_live_5WrHO8f5mvzHbrHKeTbXfTCC")
	chargeParams := &ChargeParams{
		order_no:  "88888887",
		appid:     "app_mHarHK4KajnDWDW9",
		channel:   "alipay",
		amount:    1000,
		currency:  "cny",
		client_ip: "127.0.0.1",
		subject:   "test",
		body:      "bodysample",
	}

	charge, err := chargeClient.New(chargeParams)
	return &charge
	if err != nil {
		t.Error(err)
	}
	t.Log(string(target.Id))
	t.Log(string(target.Object))
}

func TestChargeList(t *testing.T) {
	var client *ChargeClient
	client = GetChargeClient("sk_live_5WrHO8f5mvzHbrHKeTbXfTCC")
	params := &ChargeListParams{}
	target, err := client.List(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(target.Url))
}

func TestChargeGet(t *testing.T) {
	client := GetChargeClient("sk_live_5WrHO8f5mvzHbrHKeTbXfTCC")
	var id string
	id = "ch_n10ejHLibfzPPmf9i5u9uT4G"
	target, err := client.Get(id)
	if err != nil {
		t.Error(err)
	}
	t.Log(target)
}

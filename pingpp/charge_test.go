package pingpp

import (
	"testing"
)

func TestChargeNew(t *testing.T) {

	chargeClient := getChargeClient("sk_live_5WrHO8f5mvzHbrHKeTbXfTCC")
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

	charge, err := chargeClient.new(chargeParams)
	return &charge
	if err != nil {
		t.Error(err)
	}
	t.Log(string(target.Id))
	t.Log(string(target.Object))
}

func TestChargeList(t *testing.T) {
	var client *ChargeClient
	client = getChargeClient("sk_live_5WrHO8f5mvzHbrHKeTbXfTCC")
	params := &ChargeListParams{}
	target, err := client.list(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(target.Url))
}

func TestChargeGet(t *testing.T) {
	client := getChargeClient("sk_live_5WrHO8f5mvzHbrHKeTbXfTCC")
	var id string
	id = "ch_n10ejHLibfzPPmf9i5u9uT4G"
	target, err := client.get(id)
	if err != nil {
		t.Error(err)
	}
	t.Log(target)
}

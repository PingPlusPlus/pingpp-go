package pingpp

import (
	"testing"
)

func TestRedEnvelopeNew(t *testing.T) {

	redEnvelopeClient := GetRedEnvelopeClient("sk_live_bPaPC8TyvHuD8uXjfPSKGa90")
	redenvelopeParams := &redEnvelopeParams{
		order_no: "88888889",
		app: {"id":"app_9Kafv5qD0iP4jH48"},
		channel:  "wx_pub",
		amount:   10,
		currency: "cny",
		subject:  "Your Subject",
		body:     "Your Body",
		extra: {"nick_name":"Nick Name","send_name":"Send Name"},
		description: "Your Description"
	}

	redenvelope, err := redEnvelopeClient.New(redenvelopeParams)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(redenvelope.Id))
	t.Log(string(redenvelope.Object))
	t.Log(string(redenvelope.App))

}

func TestRedEnvelopeList(t *testing.T) {
	var client *RedEnvelopeClient
	client = GetRedEnvelopeClient("sk_live_bPaPC8TyvHuD8uXjfPSKGa90")
	params := &RedEnvelopeListParams{}
	target, err := client.List(params)
	if err != nil {
		t.Error(err)
	}
	t.Error(string(target.Url))
}

func TestRedEnvelopeGet(t *testing.T) {
	client := GetRedEnvelopeClient("sk_live_bPaPC8TyvHuD8uXjfPSKGa90")
	var id string
	id = "red_5iHq5CDO40eHSKCGi104qzLG"
	target, err := client.Get(id)
	if err != nil {
		t.Error(err)
	}
	t.Error(target)
}

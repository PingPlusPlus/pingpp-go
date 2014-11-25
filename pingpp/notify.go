package pingpp

import (
	"encoding/json"
	fb "github.com/huandu/facebook"
)

type NotifyCharge struct {
	Id              string
	Object          string            `facebook:"object"`
	Created         string            `facebook:"created"`
	Livemode        bool              `facebook:"livemode"`
	Paid            bool              `facebook:"paid"`
	Refunded        bool              `facebook:"refunded"`
	Order_no        string            `facebook:"order_no"`
	App             *App              `facebook:"app"`
	Channel         string            `facebook:"channel"`
	Amount          uint64            `facebook:"amount"`
	Amount_settle   uint64            `facebook:"amount_settle"`
	Amount_refunded uint64            `facebook:"amount_refunded"`
	Time_expire     string            `facebook:"time_expire"`
	Time_settle     string            `facebook:"time_settle"`
	Transaction_no  string            `facebook:"transaction_no"`
	Currency        string            `facebook:"currency"`
	Client_ip       string            `facebook:"client_ip"`
	Subject         string            `facebook:"subject"`
	Body            string            `facebook:"body"`
	Failure_code    int               `facebook:"failure_code"`
	Failure_msg     string            `facebook:"failure_msg"`
	Metadata        map[string]string `facebook:"metadata"`
	Refunds         *RefundList       `facebook:"refunds"`
	Credential      *Credential       `facebook:"credential"`
}

type App struct {
	Id                 string
	Object             string
	Created            uint64
	Display_name       string
	Notify_url         string
	Goods_type         int64
	Channels_supported []string
}
type ObjectIndentify struct {
	Object string `facebook:object`
}

func parseNotify(notifyJson string) interface{} {
	var identify ObjectIndentify
	var charge NotifyCharge
	var refund Refund
	var jsObject fb.Result
	err := json.Unmarshal([]byte(notifyJson), &identify)

	if err != nil {
		return nil
	}
	if identify.Object == "charge" {
		err2 := json.Unmarshal([]byte(notifyJson), &jsObject)
		if err2 != nil {
			decodeError := jsObject.Decode(&NotifyCharge)
			return &NotifyCharge
		} else {
			return nil
		}
	} else if identify.Object == "refund" {
		err2 := json.Unmarshal([]byte(notifyJson), &refund)
		if err2 != nil {
			return &refund
		} else {
			return nil
		}
	}
	return nil
}

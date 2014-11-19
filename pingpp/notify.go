package pingpp

import (
	"encoding/json"
)

type ObjectIndentify struct {
	object string `json:object`
}

func parseNotify(notifyJson string) interface{} {
	var identify ObjectIndentify
	var charge Charge
	var refund Refund
	err := json.Unmarshal([]byte(notifyJson), &identify)

	if err != nil {
		return nil
	}
	if identify.object == "charge" {
		err2 := json.Unmarshal([]byte(notifyJson), &charge)
		if err2 != nil {
			return &charge
		} else {
			return nil
		}
	} else if identify.object == "refund" {
		err2 := json.Unmarshal([]byte(notifyJson), &refund)
		if err2 != nil {
			return &refund
		} else {
			return nil
		}
	}
	return nil
}

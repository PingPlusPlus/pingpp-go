package pingpp

import (
	"encoding/json"
)


type ObjectIndentify struct {
	object string `json:object`
}

func parseNotify(notifyJson string){
	var identify ObjectIndentify
	var charge Charge
	var refund Refund
	err := json.Unmarshal(notifyJson, &identify)
	if err != nil {
		return nil
	}
	if identify.object == "charge" {
		return json.Unmarshal(notifyJson, &charge)
	}
	else if identify.object == "refund" {
		return json.Unmarshal(notifyJson, &refund)
	}
	return nil
}
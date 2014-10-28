package pingpp

import (
	"strconv"
	"testing"
)

func TestRefundNew(t *testing.T) {
	client := getRefundClient("sk_live_5WrHO8f5mvzHbrHKeTbXfTCC")
	refundParams := &RefundParams{
		Amount:      1,
		Description: "Lekton Test",
	}
	t.Logf("Amount: %v\n", strconv.FormatInt(refundParams.Amount, 10))
	target, err := client.new(refundParams, "ch_r1KCmPvXPWn9qvPKW9Pu5OCC")
	if err != nil {
		t.Error(err)
	}
	t.Error(string(target.ID))
	t.Error(string(target.Object))
}

func TestRefundList(t *testing.T) {
	client := getRefundClient("sk_live_5WrHO8f5mvzHbrHKeTbXfTCC")
	target, err := client.list("ch_r1KCmPvXPWn9qvPKW9Pu5OCC", 10, "", "")
	if err != nil {
		t.Error(err)
	}
	t.Error(string(target.Url))
}

func TestRefundGet(t *testing.T) {
	client := getRefundClient("sk_live_5WrHO8f5mvzHbrHKeTbXfTCC")
	id := "ch_r1KCmPvXPWn9qvPKW9Pu5OCC"
	target, err := client.get(id, "re_WbvXLOynb1O4ifzHuHDSm1KK")
	if err != nil {
		t.Error(err)
	}
	t.Error(target)
}

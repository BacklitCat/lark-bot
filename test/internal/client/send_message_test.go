package client_test

import (
	"larkbot/internal/client"
	"strconv"
	"testing"
	"time"
)

func TestSendTextMsgSimple(t *testing.T) {
	// tian ou_5490f162c8326c0f14782b33f86cc669
	// chang ou_8f578425df90530809910b0f20f83478
	err := client.SendTextMsgSimple(
		"ou_5490f162c8326c0f14782b33f86cc669",
		"[TEST] text msg from tianqi's dev linux, unix time: "+
			strconv.Itoa(int(time.Now().Unix())))
	if err != nil {
		t.Error(err)
	}
}

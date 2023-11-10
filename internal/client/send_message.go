package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"larkbot/internal/config"
	"larkbot/larktype"
)

func SendTextMsgSimple(receiveID, msgContent string) error {
	req := larktype.MessagesReq{
		ReceiveID: receiveID,
		MsgType:   "text",
		Content:   msgContent,
		UUID:      "",
	}

	reqBody, err := json.Marshal(&req)
	if err != nil {
		return err
	}

	if len(reqBody) == 0 {
		return errors.New("HTTP playload is empty")
	}

	header, err := GetTenentAccessTokenHeader()
	if err != nil {
		return err
	}

	httpRespBody, err := doHTTP(reqBody, config.LarkAPI.MESSAGES.Method, config.LarkAPI.MESSAGES.Url+"?receive_id_type=open_id", header)

	if err != nil {
		return err
	}

	var resp larktype.MessagesResp
	if err = json.Unmarshal(httpRespBody, &resp); err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf("lark say, code:%d msg: %s", resp.Code, resp.Msg)
	}
	return nil
}

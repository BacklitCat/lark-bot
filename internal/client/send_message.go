package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"larkbot/internal/config"
	"larkbot/larktype"
)

// SendMsgSimple
// https://open.feishu.cn/document/server-docs/im-v1/message/create?appId=cli_a5c372574e18100b
func SendMsgSimple(receiveID, receiveIDType, msgType, msgContent string) error {
	req := larktype.MessagesReq{
		ReceiveID: receiveID,
		MsgType:   msgType,
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

	respBody, err := doHttpDefaultClient(
		reqBody,
		config.LarkAPI.MESSAGES.Method,
		fmt.Sprintf("%s?receive_id_type=%s", config.LarkAPI.MESSAGES.Url, receiveIDType),
		header)

	if err != nil {
		return err
	}

	var resp larktype.MessagesResp
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf("lark say, code:%d msg: %s", resp.Code, resp.Msg)
	}
	return nil
}

func SendTextMsgSimple(receiveID, msgStr string) error {
	return SendMsgSimple(receiveID, "open_id", "text",
		fmt.Sprintf("{\"text\":\"%s\"}", msgStr))
}

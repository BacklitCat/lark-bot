package handler

import (
	"context"
	"fmt"
	"larkbot/internal/client"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/sirupsen/logrus"
)

// RepeatMachineHandler
func RepeatMachineHandler(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	senderOpenId := *event.Event.Sender.SenderId.OpenId
	messageContent := *event.Event.Message.Content
	logrus.WithFields(logrus.Fields{
		"Sender.SenderId.OpenId":   senderOpenId,
		"Message.Content":          messageContent,
		"larkcore.Prettify(event)": larkcore.Prettify(event),
		"event.RequestId()":        event.RequestId(),
	})

	err := client.SendTextMsgContentSimple(senderOpenId, *event.Event.Message.Content)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

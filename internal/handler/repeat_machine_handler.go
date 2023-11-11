package handler

import (
	"context"
	"larkbot/internal/client"
	"larkbot/internal/logger"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/sirupsen/logrus"
)

// RepeatMachineHandler
func RepeatMachineHandler(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	senderOpenId := *event.Event.Sender.SenderId.OpenId
	messageContent := *event.Event.Message.Content

	_ = client.SendTextMsgContentSimple(senderOpenId, *event.Event.Message.Content)

	// 测试日志
	logger.EventLogger.WithFields(logrus.Fields{
		"Sender.SenderId.OpenId": senderOpenId,
		"Message.Content":        messageContent,
		// "larkcore.Prettify(event)": larkcore.Prettify(event),
		"event.RequestId()": event.RequestId(),
	}).Info()

	return nil
}

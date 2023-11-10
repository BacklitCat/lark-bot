package main

import (
	"context"
	"fmt"
	"larkbot/internal/client"
	"larkbot/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	sdkginext "github.com/larksuite/oapi-sdk-gin"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func main() {

	// var client = lark.NewClient(config.BotConf.Lark.AppID, config.BotConf.Lark.AppSecret,
	// 	lark.WithLogLevel(larkcore.LogLevelDebug),
	// 	lark.WithReqTimeout(3*time.Second),
	// 	lark.WithEnableTokenCache(true),
	// 	lark.WithHelpdeskCredential("id", "token"),
	// 	lark.WithHttpClient(http.DefaultClient))

	// 注册消息处理器
	handler := dispatcher.NewEventDispatcher(config.Bot.Lark.VerificationToken, config.Bot.Lark.EncryptKey).
		// 机器人接收到用户发送的消息后触发此事件
		OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
			fmt.Println(larkcore.Prettify(event))
			fmt.Println(event.RequestId())

			// 临时做个复读机
			fmt.Println(*event.Event.Sender.SenderId.OpenId, *event.Event.Message.Content)
			err := client.SendTextMsgSimple(*event.Event.Sender.SenderId.OpenId, *event.Event.Message.Content)
			if err != nil {
				fmt.Println(err)
			}

			return nil
		}).
		// 用户阅读机器人发送的单聊消息后触发此事件
		OnP2MessageReadV1(func(ctx context.Context, event *larkim.P2MessageReadV1) error {
			fmt.Println(larkcore.Prettify(event))
			fmt.Println(event.RequestId())
			return nil
		})

	// 创建卡片行为处理器
	cardHandler := larkcard.NewCardActionHandler(config.Bot.Lark.VerificationToken, config.Bot.Lark.EncryptKey,
		func(ctx context.Context, cardAction *larkcard.CardAction) (interface{}, error) {
			fmt.Println(larkcore.Prettify(cardAction))

			// 返回卡片消息
			//return getCard(), nil

			//custom resp
			//return getCustomResp(),nil

			// 无返回值
			return nil, nil
		})

	g := gin.Default()

	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "lark_bot")
	})

	g.POST("/webhook/event", sdkginext.NewEventHandlerFunc(handler))
	g.POST("/webhook/card", sdkginext.NewCardActionHandlerFunc(cardHandler))

	err := g.Run(":" + config.Bot.Server.Port)
	if err != nil {
		panic(err)
	}
}

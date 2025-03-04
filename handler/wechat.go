package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

var officialAccount *officialaccount.OfficialAccount

func SetupWechatHandler(cfg *config.Config) {
	wc := wechat.NewWechat()
	officialAccount = wc.GetOfficialAccount(cfg)
}

func WechatHandler(c *gin.Context) {
	server := officialAccount.GetServer(c.Request, c.Writer)

	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		switch msg.Event {
		case message.EventSubscribe:
			test :=&message.Reply{
				MsgType: message.MsgTypeText,
				MsgData: message.NewText(fmt.Sprintf("欢迎关注公众号！EventKey: %s", msg.EventKey)),
			}
			println(test.MsgData)
			return test
		case message.EventScan:
			return &message.Reply{
				MsgType: message.MsgTypeText,
				MsgData: message.NewText("您已经关注该账号"),
			}
		default:
			return &message.Reply{
				MsgType: message.MsgTypeText,
				MsgData: message.NewText("你好"),
			}
		}
	})

	if err := server.Serve(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	server.Send()
}

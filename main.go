package main

import (
	"phidialam/wx-oa-demo/config"
	"phidialam/wx-oa-demo/handler"
	"phidialam/wx-oa-demo/server"

	"github.com/silenceper/wechat/v2/cache"
	oaConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	oaCache := cache.NewMemcache()
	wxConfig := &oaConfig.Config{
		AppID:          cfg.Wechat.AppID,
		AppSecret:      cfg.Wechat.AppSecret,
		Token:          cfg.Wechat.Token,
		EncodingAESKey: cfg.Wechat.EncodingAESKey,
		Cache: oaCache,
	}

	handler.SetupWechatHandler(wxConfig)

	router := server.SetupRouter()
	router.Any("/wechat", handler.WechatHandler)

	if err := server.Start(router, cfg.Server.Port); err != nil {
		panic(err)
	}
}

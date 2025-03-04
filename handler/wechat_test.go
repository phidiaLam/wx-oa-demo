package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount/config"
)

func TestSubscribeEvent(t *testing.T) {
	// 初始化测试配置
	oaCache := cache.NewMemcache()
	cfg := &config.Config{
		Token: "test_token",
		Cache: oaCache,
	}
	SetupWechatHandler(cfg)

	// 创建测试上下文
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// 构造请求参数
	params := url.Values{}
	params.Set("signature", "dc24bd81e824da6148aab5bddce98cb9de93533b") // 这里需要根据实际 token 计算
	params.Set("timestamp", "123456789")
	params.Set("nonce", "nonce123")

	// 构造 POST 请求体
	xmlBody := `<xml>
	  <ToUserName><![CDATA[toUser]]></ToUserName>
	  <FromUserName><![CDATA[FromUser]]></FromUserName>
	  <CreateTime>1348831860</CreateTime>
	  <MsgType><![CDATA[event]]></MsgType>
	  <Event><![CDATA[subscribe]]></Event>
	  <EventKey><![CDATA[qrscene_123]]></EventKey>
	</xml>`

	// 设置请求
	c.Request = httptest.NewRequest("POST", "/wechat?"+params.Encode(), strings.NewReader(xmlBody))
	c.Request.Header.Set("Content-Type", "application/xml")

	// 执行处理
	WechatHandler(c)

	// 验证响应
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	expected := "欢迎关注公众号！EventKey: qrscene_123"
	if !strings.Contains(w.Body.String(), expected) {
		t.Errorf("Expected response to contain %q, got %q", expected, w.Body.String())
	}
}

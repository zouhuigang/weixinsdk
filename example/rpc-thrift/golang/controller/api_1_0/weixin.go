package api_1_0

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"weixinsdk/example/rpc-thrift/golang/utils"

	"github.com/labstack/echo"
)

type WeixinApi struct{}

const (
	token = "weixin" //设置token
)

var weixinApi = WeixinApi{}

//注册路由
func (this *WeixinApi) RegisterRoute(g *echo.Group) {
	g.Any("/weixin/callback", this.wx_callback)
}

func (this *WeixinApi) wx_callback(ctx echo.Context) error {

	method := ctx.Request().Method
	fmt.Println(method)
	if method == "GET" {
		timestamp := ctx.FormValue("timestamp")
		nonce := ctx.FormValue("nonce")
		signature := ctx.FormValue("signature")
		echostr := ctx.FormValue("echostr")

		wxServerClient := utils.GetWxServerClient()
		wx, err := wxServerClient.IsWeixinServer(token, echostr, signature, timestamp, nonce)
		if err != nil {

		}
		if !wx.IsServer {
			return ctx.String(http.StatusOK, "验证失败")
		}
		//验证成功，原样返回echostr
		return ctx.String(http.StatusOK, wx.Echostr)

	} else if method == "POST" {

		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			return ctx.String(http.StatusOK, "")
		}

		if len(body) == 0 {
			return ctx.String(http.StatusOK, "")
		}
		wechatReply, err := logic.DefaultWechat.AutoReply(body)
	if err != nil {
		return ctx.String(http.StatusOK, "")
	}

	return ctx.XML(http.StatusOK, wechatReply)
	}
	return ctx.String(http.StatusOK, "提交方式错误")
}

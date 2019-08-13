package api_1_0

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"weixinsdk/example/rpc-thrift/golang/utils"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/weixin/service" //注意导入Thrift生成的接口包

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

/*
一旦遇到以下情况，微信都会在公众号会话中，向用户下发系统提示“该公众号暂时无法提供服务，请稍后再试”：

1、开发者在5秒内未回复任何内容
2、开发者回复了异常数据，比如JSON数据等
另外，请注意，回复图片（不支持gif动图）等多媒体消息时需要预先通过素材管理接口上传临时素材到微信服务器，可以使用素材管理中的临时素材，也可以使用永久素材。
*/
func (this *WeixinApi) wx_callback(ctx echo.Context) error {

	method := ctx.Request().Method
	wxServerClient := utils.GetWxServerClient()
	if method == "GET" {
		timestamp := ctx.FormValue("timestamp")
		nonce := ctx.FormValue("nonce")
		signature := ctx.FormValue("signature")
		echostr := ctx.FormValue("echostr")

		wx, err := wxServerClient.IsWeixinServer(token, echostr, signature, timestamp, nonce)
		if err != nil {

		}
		if !wx.IsServer {
			return ctx.String(http.StatusOK, "验证失败")
		}
		//验证成功，原样返回echostr
		return ctx.String(http.StatusOK, wx.Echostr)

	} else if method == "POST" {
		var responeXmlStr string = "success"
		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			return ctx.String(http.StatusOK, "")
		}

		if len(body) == 0 {
			return ctx.String(http.StatusOK, "")
		}
		mixMessage, err := wxServerClient.ParseTemplateToMixedMessages(string(body))
		//SendTmplateMessage
		if err != nil {
			return ctx.String(http.StatusOK, "消息解析失败")
		}
		if mixMessage.ResponeMessageType == "text" {
			var contentText string = *mixMessage.ResponeMessage.Content
			if contentText == "模板1" {
				templateMsg := z_weixin_service.NewTemplateMsgData()
				tempData := z_weixin_service.NewTemplateData()

				tempData.First = &z_weixin_service.KeyWordData{Value: "测试模板消息"}
				tempData.Keyword1 = &z_weixin_service.KeyWordData{Value: "大家记得买票啊"}
				tempData.Keyword2 = &z_weixin_service.KeyWordData{Value: "马上就要放假了，大家记得买回家的票啊"}
				tempData.Keyword3 = &z_weixin_service.KeyWordData{Value: "2018-12-30 15:59"}
				tempData.Keyword4 = &z_weixin_service.KeyWordData{Value: "派大星"}
				tempData.Keyword5 = &z_weixin_service.KeyWordData{Value: "记得按时完成"}

				templateMsg.Data = tempData
				templateMsg.FormID = mixMessage.ResponeMessage.ToUserName
				templateMsg.Touser = mixMessage.ResponeMessage.FromUserName
				templateMsg.TemplateID = "byP6PqRaW34ZSxKXE4eZOX1TCBSbZiWZiFfoYJPxN9Y"
				respone, err := wxServerClient.SendTmplateMessage(templateMsg)
				fmt.Println(respone, err)
			} else {
				fromUserName := mixMessage.ResponeMessage.ToUserName
				toUserName := mixMessage.ResponeMessage.FromUserName
				responeXmlStr, _ = wxServerClient.GetTextXml(fromUserName, toUserName, contentText)

			}

		} else if mixMessage.ResponeMessageType == "image" {

		}

		return ctx.String(http.StatusOK, responeXmlStr)
	}
	return ctx.String(http.StatusOK, "提交方式错误")
}

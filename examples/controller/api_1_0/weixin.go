package api_1_0

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	zcore "weixinsdk/src/core/service"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/tencent/weixin/service" //注意导入Thrift生成的接口包

	"github.com/labstack/echo"
)

type WeixinApi struct{}

const (
	token  = "weixin" //设置token
	WEBURL = "http://4ggmu7.natappfree.cc/"
)

var weixinApi = WeixinApi{}

//注册路由
func (this *WeixinApi) RegisterRoute(g *echo.Group) {
	g.Any("/weixin/callback", this.wx_callback)
	g.Any("/weixin/snsapi_base/page1", this.snsapi_base_page1)
	g.Any("/weixin/snsapi_base/page2", this.snsapi_base_page2)

	g.Any("/weixin/snsapi_userinfo/page1", this.snsapi_userinfo_page1)
	g.Any("/weixin/snsapi_userinfo/page2", this.snsapi_userinfo_page2)
}

/*
一旦遇到以下情况，微信都会在公众号会话中，向用户下发系统提示“该公众号暂时无法提供服务，请稍后再试”：

1、开发者在5秒内未回复任何内容
2、开发者回复了异常数据，比如JSON数据等
另外，请注意，回复图片（不支持gif动图）等多媒体消息时需要预先通过素材管理接口上传临时素材到微信服务器，可以使用素材管理中的临时素材，也可以使用永久素材。
*/
func (this *WeixinApi) wx_callback(ctx echo.Context) error {
	handler := &zcore.WxServiceThrift{}

	method := ctx.Request().Method
	if method == "GET" {
		timestamp := ctx.FormValue("timestamp")
		nonce := ctx.FormValue("nonce")
		signature := ctx.FormValue("signature")
		echostr := ctx.FormValue("echostr")

		wx, err := handler.IsWeixinServer(token, echostr, signature, timestamp, nonce)
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
		mixMessage, err := handler.ParseTemplateToMixedMessages(string(body))
		//SendTmplateMessage
		if err != nil {
			return ctx.String(http.StatusOK, "消息解析失败")
		}

		fromUserName := mixMessage.ResponeMessage.ToUserName
		toUserName := mixMessage.ResponeMessage.FromUserName
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
				respone, err := handler.SendTmplateMessage(templateMsg)
				fmt.Println(respone, err)
			} else if contentText == "素材数量" {
				resCoun, _ := handler.MaterialCount()
				contentText = fmt.Sprintf("【永久素材数量】语音:%d-视频:%d-图片:%d-图文:%d:", resCoun.VoiceCount, resCoun.VideoCount, resCoun.ImageCount, resCoun.NewsCount_)
				responeXmlStr, _ = handler.GetTextXml(fromUserName, toUserName, contentText)
			} else if contentText == "上传素材" {
				resU, _ := handler.UpImage(`image`, `D:\多可德更新\陪诊关注推送图\多可德推图1.jpg`)
				fmt.Println(resU)
				contentText = fmt.Sprintf("【上传图片】media_id:%s", resU.MediaID)
				responeXmlStr, _ = handler.GetTextXml(fromUserName, toUserName, contentText)
			} else if contentText == "图片1" {
				imageData := z_weixin_service.NewImageData()
				imageData.MediaId = `iZfVswoy_wsLl4zaxcs2j7Y7px49j9JBvyFQ-xsJEQY`

				autoReply := z_weixin_service.NewAutoReplyData()
				autoReply.FromUserName = fromUserName
				autoReply.ToUserName = toUserName
				autoReply.MsgType = `image`
				autoReply.Image = imageData
				responeXmlStr, _ = handler.GetAutoReplyXml(autoReply)
				fmt.Println(responeXmlStr)
			} else if contentText == "图文1" {
				aList := make([]*z_weixin_service.ArticlesData, 0)
				articlesData := z_weixin_service.NewArticlesData()
				articlesData.Title = `这是标题`
				articlesData.Description = `这是描述`
				articlesData.URL = `https://www.baidu.com`
				articlesData.PicUrl = `https://cdn-oss.yyang.net.cn/static/vue_image/huize_about.jpg`
				aList = append(aList, articlesData)

				articlesData2 := z_weixin_service.NewArticlesData()
				articlesData2.Title = `这是标题2`
				articlesData2.Description = `这是描述2`
				articlesData2.URL = `https://www.baidu.com`
				articlesData2.PicUrl = `https://cdn-oss.yyang.net.cn/static/vue_image/huize_about.jpg`
				aList = append(aList, articlesData2)

				autoReply := z_weixin_service.NewAutoReplyData()
				autoReply.FromUserName = fromUserName
				autoReply.ToUserName = toUserName
				autoReply.MsgType = `news`
				autoReply.ArticleCount = 2
				autoReply.Articles = aList
				responeXmlStr, _ = handler.GetAutoReplyXml(autoReply)
				fmt.Println(responeXmlStr)
			} else {
				responeXmlStr, _ = handler.GetTextXml(fromUserName, toUserName, contentText)
			}

		} else if mixMessage.ResponeMessageType == "image" { //转发给客服
			responeXmlStr, _ = handler.TransferCustomerService(fromUserName, toUserName, "")
			fmt.Println("微信", responeXmlStr)

		}

		return ctx.String(http.StatusOK, responeXmlStr)
	}
	return ctx.String(http.StatusOK, "提交方式错误")
}

//微信回调
func (this *WeixinApi) snsapi_base_page1(ctx echo.Context) error {
	handler := &zcore.WxServiceThrift{}

	//构造跳转链接
	redirectURL := fmt.Sprintf("%sweixin/snsapi_base/page2", WEBURL)
	scope := "snsapi_base"
	authData, _ := handler.AuthCodeURL(redirectURL, scope)

	fmt.Println("可以将state保存在cookie中,在下一个页面验证url中的state和保存的是否一致", authData.State)
	//ctx.Redirect(301, "/welcome")/
	return ctx.Redirect(301, authData.URL)
}

func (this *WeixinApi) snsapi_base_page2(ctx echo.Context) error {
	handler := &zcore.WxServiceThrift{}

	var code string = ctx.FormValue("code")
	if code == "" {
		return ctx.String(http.StatusOK, "客户禁止授权")
	}
	userinfo, err := handler.GetUserInfoBySnsapiBase(code)
	if err != nil {
		return ctx.String(http.StatusOK, "无感授权失败"+err.Error())
	}

	buf, err := json.MarshalIndent(userinfo, "", "    ") //格式化编码
	if err != nil {
		return ctx.String(http.StatusOK, "json解析失败")
	}
	return ctx.String(http.StatusOK, string(buf))
}

func (this *WeixinApi) snsapi_userinfo_page1(ctx echo.Context) error {
	handler := &zcore.WxServiceThrift{}

	//构造跳转链接
	redirectURL := fmt.Sprintf("%sweixin/snsapi_base/page2", WEBURL)
	scope := "snsapi_userinfo"
	authData, _ := handler.AuthCodeURL(redirectURL, scope)

	fmt.Println("可以将state保存在cookie中,在下一个页面验证url中的state和保存的是否一致", authData.State)
	//ctx.Redirect(301, "/welcome")/
	return ctx.Redirect(301, authData.URL)
}

func (this *WeixinApi) snsapi_userinfo_page2(ctx echo.Context) error {
	handler := &zcore.WxServiceThrift{}

	var code string = ctx.FormValue("code")
	fmt.Println("code", code)
	if code == "" {
		return ctx.String(http.StatusOK, "客户禁止授权")
	}

	userinfo, err := handler.GetUserInfoBySnsapiUserinfo(code)
	if err != nil {
		return ctx.String(http.StatusOK, "用户授权失败"+err.Error())
	}

	buf, err := json.MarshalIndent(userinfo, "", "    ") //格式化编码
	if err != nil {
		return ctx.String(http.StatusOK, "json解析失败")
	}
	return ctx.String(http.StatusOK, string(buf))
}

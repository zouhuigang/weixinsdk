package service

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strings"
	zcache "weixinsdk/src/cache"
	zconfig "weixinsdk/src/config"
	"weixinsdk/src/logger"
	"weixinsdk/src/structure"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/tencent/weixin/service" //注意导入Thrift生成的接口包

	"github.com/zouhuigang/package/zhttp"
	"github.com/zouhuigang/package/ztime"
)

//https://www.jianshu.com/p/8a21ce49be20
//https://github.com/bigwhite/experiments/blob/master/wechat_examples/public/2-recvtextmsg/recvtextmsg_unencrypt.go
func makeSignature(token, timestamp, nonce string) string { //本地计算signature
	si := []string{token, timestamp, nonce}
	sort.Strings(si)            //字典序排序
	str := strings.Join(si, "") //组合字符串
	s := sha1.New()             //返回一个新的使用SHA1校验的hash.Hash接口
	io.WriteString(s, str)      //WriteString函数将字符串数组str中的内容写入到s中
	return fmt.Sprintf("%x", s.Sum(nil))
}

/*
当普通微信用户向公众账号发消息时，微信服务器将POST消息的XML数据包到开发者填写的URL上。
验证消息来源于微信
token:是固定的，例如:weixin
signature, timestamp, nonce,echostr都是微信服务器通过get传过来的

timestamp := ctx.FormValue("timestamp")
nonce := ctx.FormValue("nonce")
signature := ctx.FormValue("signature")
echostr := ctx.FormValue("echostr")
*/
func (this *WxServiceThrift) IsWeixinServer(token, echostr, signature, timestamp, nonce string) (*z_weixin_service.IsWeixinServerData, error) {

	data := z_weixin_service.NewIsWeixinServerData()
	signatureGen := makeSignature(token, timestamp, nonce)

	m_log := fmt.Sprintf("IsWeixinServer get data:%s-%s-%s-%s-%s", token, echostr, signature, timestamp, nonce)
	logger.MyLogger.Debug(m_log)
	if signatureGen != signature {
		data.IsServer = false
		data.Echostr = ""

	} else {
		data.IsServer = true
		data.Echostr = echostr
	}

	return data, nil
}

//解析模板消息,为混合消息
/*
微信服务器在五秒内收不到响应会断掉连接，并且重新发起请求，总共重试三次。假如服务器无法保证在五秒内处理并回复，

可以直接回复空串，微信服务器不会对此作任何处理，并且不会发起重试。详情请见“发送消息-被动回复消息”。
*/
func (this *WxServiceThrift) ParseTemplateToMixedMessages(body string) (*z_weixin_service.ParseTemplateToMixedMessagesData, error) {
	data := z_weixin_service.NewParseTemplateToMixedMessagesData()
	logger.MyLogger.Debug("Receiving messages from Wechat:" + body)

	//解析消息
	msg := &z_weixin_service.MixedMessage{}
	err := xml.Unmarshal([]byte(body), &msg)
	if err != nil {
		logger.MyLogger.Error("ParseTemplateMessages" + err.Error())
		return data, err
	}
	data.ResponeMessageType = msg.MsgType
	data.ResponeMessage = msg

	return data, nil

}

//发送模板消息
func (this *WxServiceThrift) SendTmplateMessage(msg *z_weixin_service.TemplateMsgData) (*z_weixin_service.SendTemplateResponseData, error) {
	response := z_weixin_service.NewSendTemplateResponseData()
	access_token, err := zcache.GetAccessToken()
	if err != nil {
		logger.MyLogger.Error("SendTmplateMessage" + err.Error())
		return nil, err
	}

	lg := fmt.Sprintf("Receiving messages to Wechat:%v", msg)
	logger.MyLogger.Debug(lg)

	requrl := fmt.Sprintf("%s?access_token=%s", zconfig.SERVICE_APIURL_SEND_TEMPLATE, access_token)
	err = zhttp.POSTWithUnmarshal(requrl, msg, response)

	return response, err
}

// 文本消息
// type Text struct {
// 	XMLName struct{} `xml:"xml" json:"-"`
// 	core.MsgHeader
// 	Content string `xml:"Content" json:"Content"` // 回复的消息内容(换行: 在content中能够换行, 微信客户端支持换行显示)
// }

//被动回复消息https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140543
//text.FromUserName, text.ToUserName, text.CreateTime, text.Content
func (this *WxServiceThrift) GetTextXml(fromUserName, toUserName, content string) (string, error) {
	var msg structure.WxMsgTxt
	msg.MsgType = structure.WeMsgTypeText
	msg.FromUserName = fromUserName
	msg.ToUserName = toUserName
	msg.Content = content
	msg.CreateTime = ztime.NowTimeStamp()

	data, err := xml.MarshalIndent(&msg, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil

}

//回复图片

//发送消息
// func (this *WxServiceThrift) GetText(msg *structure.MixedMessage) *structure.WxMsgTxt {
// 	return &structure.WxMsgTxt{
// 		CommonMessageHeader: msg.CommonMessageHeader,
// 		Content:             msg.Content,
// 	}
// }

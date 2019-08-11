package service

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"

	"strings"
	"weixinsdk/src/logger"
	"weixinsdk/src/structure"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/weixin/service" //注意导入Thrift生成的接口包
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

//解析模板消息

//发送消息

func GetText(msg *structure.MixedMessage) *structure.WxMsgTxt {
	return &structure.WxMsgTxt{
		CommonMessageHeader: msg.CommonMessageHeader,
		MsgId:               msg.MsgId,
		Content:             msg.Content,
	}
}

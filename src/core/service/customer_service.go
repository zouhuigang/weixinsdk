package service

import (
	"encoding/xml"
	"weixinsdk/src/structure"

	"github.com/zouhuigang/package/ztime"
)

/*
客服消息
如果公众号处于开发模式，普通微信用户向公众号发消息时，微信服务器会先将消息POST到开发者填写的url上，
如果希望将消息转发到客服系统，则需要开发者在响应包中返回MsgType为transfer_customer_service的消息，
微信服务器收到响应后会把当次发送的消息转发至客服系统。
您也可以在返回transfer_customer_service消息时，在XML中附上TransInfo信息指定分配给某个客服帐号。
*/

//转发给客服
func (this *WxServiceThrift) TransferCustomerService(fromUserName, toUserName, kfAccount string) (string, error) {
	var msg structure.WxMsgTransferCustomerService
	msg.MsgType = structure.WeTransferCustomerService
	msg.FromUserName = fromUserName
	msg.ToUserName = toUserName
	msg.KfAccount = kfAccount
	msg.CreateTime = ztime.NowTimeStamp()

	data, err := xml.MarshalIndent(&msg, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil

}

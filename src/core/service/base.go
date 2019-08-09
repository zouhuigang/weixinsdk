package service

import (
	zcache "weixinsdk/src/cache"
	//注意导入Thrift生成的接口包
)

func (this *WxServiceThrift) GetJsapiTicket() (string, error) {
	jsapi_ticket, err := zcache.GetJsapiTicket()
	return jsapi_ticket, err
}

func (this *WxServiceThrift) GetAccessToken() (string, error) {
	token, err := zcache.GetAccessToken()
	return token, err
}

package service

import (
	"fmt"
	zcache "weixinsdk/src/cache"
	zconfig "weixinsdk/src/config"
	zutils "weixinsdk/src/utils"

	"github.com/zouhuigang/package/ztime"
)

//jssdk 扫一扫，分享等
func Jsapi_sign(url string) map[string]interface{} {

	tmp := map[string]string{}

	//参数
	m_AppID := zconfig.CFG.MustValue("service", "AppID", "")
	jsapi_ticket := zcache.GetJsapiTicket()
	timestamp := fmt.Sprintf("%d", ztime.NowTimeStamp())
	noncestr := zutils.NonceStr()

	//获取sign
	tmp["jsapi_ticket"] = jsapi_ticket
	tmp["noncestr"] = noncestr
	tmp["timestamp"] = timestamp
	tmp["url"] = url
	sign := zutils.Signature(tmp)

	//输出给前端api
	data := map[string]interface{}{}
	data["jsapi_ticket"] = jsapi_ticket
	data["timestamp"] = timestamp
	data["noncestr"] = noncestr
	data["url"] = url
	data["sign"] = sign
	data["appid"] = m_AppID
	return data
}

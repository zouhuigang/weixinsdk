package service

// import (
// 	"fmt"
// 	zcache "weixinsdk/src/cache"
// 	zconfig "weixinsdk/src/config"
// 	z_weixin_service "weixinsdk/src/thrift_file/gen-go/tencent/weixin/service"
// 	zutils "weixinsdk/src/utils"

// 	"github.com/zouhuigang/package/ztime"
// )

// /*
// jsapi微信支付
// */
// func (this *WxServiceThrift) JsApiPay() (*z_weixin_service.JsapiSignData, error) {
// 	//输出给前端api
// 	res := z_weixin_service.NewJsapiSignData()
// 	tmp := map[string]string{}

// 	//参数
// 	m_AppID := zconfig.CFG.MustValue("service", "AppID", "")
// 	jsapi_ticket, err := zcache.GetJsapiTicket()
// 	if err != nil {
// 		return res, err
// 	}
// 	timestampInt64 := ztime.NowTimeStamp()
// 	timestamp := fmt.Sprintf("%d", timestampInt64)
// 	noncestr := zutils.NonceStr()
// 	m_package := "prepay_id=" + prepay_id

// 	//获取sign
// 	tmp["jsapi_ticket"] = jsapi_ticket
// 	tmp["noncestr"] = noncestr
// 	tmp["timestamp"] = timestamp
// 	sign := zutils.Signature(tmp)

// 	res.JsapiTicket = jsapi_ticket
// 	res.Timestamp = timestampInt64
// 	res.Noncestr = noncestr
// 	res.Sign = sign
// 	res.Package = m_package
// 	res.Appid = m_AppID

// 	return res, nil
// }

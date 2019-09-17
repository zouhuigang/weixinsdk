package service

import (
	"fmt"
	"strings"
	zconfig "weixinsdk/src/config"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/tencent/weixin/service"
	zutils "weixinsdk/src/utils"

	"github.com/zouhuigang/package/ztime"
)

// jsapi微信支付
/*{"ReturnCode":"SUCCESS","ReturnMsg":"OK",
"Appid":"wxadd472a86212b893","MchId":"1555438221",
"DeviceInfo":"",
"NonceStr":"s6d23YNXvq6f3q49",
"Sign":"383CA20BAA4431283392ED6FF9A88F37",
"ResultCode":"SUCCESS","ErrCode":"",
"ErrCodeDes":"","TradeType":"JSAPI",
"PrepayId":"wx16120616457270967d9e47cd1220432900",
"CodeUrl":""}*/
func (this *WxServiceThrift) GetJsApiParameters(unifiedOrderResult *z_weixin_service.UnifiedOrderResponse) (*z_weixin_service.JsApiParameters, error) {
	//输出给前端api
	res := z_weixin_service.NewJsApiParameters()

	//参数
	m_AppID := unifiedOrderResult.Appid
	timestampInt64 := ztime.NowTimeStamp()
	timestamp := fmt.Sprintf("%d", timestampInt64)
	noncestr := zutils.NonceStr()
	m_package := "prepay_id=" + unifiedOrderResult.PrepayId
	m_MchKey := zconfig.CFG.MustValue("service", "MchKey", "")
	//获取sign
	paySign := strings.ToUpper(zutils.MD5("appId=" + m_AppID + "&nonceStr=" + noncestr + "&package=" + m_package + "&signType=MD5&timeStamp=" + timestamp + "&key=" + m_MchKey))

	res.TimeStamp = timestampInt64
	res.NonceStr = noncestr
	res.PaySign = paySign
	res.PrepayId = unifiedOrderResult.PrepayId

	//额外参数
	// tmp := map[string]string{}
	//获取sign
	// tmp["jsapi_ticket"] = jsapi_ticket
	// tmp["noncestr"] = noncestr
	// tmp["timestamp"] = timestamp
	// sign := zutils.Signature(tmp)

	res.AppId = m_AppID
	res.Sign = paySign //暂时先用一样的

	return res, nil
}

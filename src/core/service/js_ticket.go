package service

import (
	"fmt"
	zcache "weixinsdk/src/cache"
	zconfig "weixinsdk/src/config"
	zutils "weixinsdk/src/utils"

	z_weixin_service "weixinsdk/src/thrift_file/gen-go/weixin/service" //注意导入Thrift生成的接口包

	"github.com/zouhuigang/package/ztime"
)

//jssdk 扫一扫，分享等
//// 每个方法除了定义的返回值之外还要返回一个error，包括定义成void的方法。自定义类型会在名字之后加一条下划线
// 暂时用不到context，所以忽略
func (this *WxServiceThrift) JsapiSign(url string) (*z_weixin_service.JsapiSignData, error) {
	//输出给前端api
	res := z_weixin_service.NewJsapiSignData()
	tmp := map[string]string{}

	//参数
	m_AppID := zconfig.CFG.MustValue("service", "AppID", "")
	jsapi_ticket, err := zcache.GetJsapiTicket()
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}
	fmt.Println("jsapi_ticket==========", jsapi_ticket)
	timestampInt64 := ztime.NowTimeStamp()
	timestamp := fmt.Sprintf("%d", timestampInt64)
	noncestr := zutils.NonceStr()

	//获取sign
	tmp["jsapi_ticket"] = jsapi_ticket
	tmp["noncestr"] = noncestr
	tmp["timestamp"] = timestamp
	tmp["url"] = url
	sign := zutils.Signature(tmp)

	res.JsapiTicket = jsapi_ticket
	res.Timestamp = timestampInt64
	res.Noncestr = noncestr
	res.URL = url
	res.Sign = sign
	res.Appid = m_AppID

	return res, nil
}

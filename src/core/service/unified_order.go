package service

import (
	"errors"
	"fmt"
	zconfig "weixinsdk/src/config"
	"weixinsdk/src/logger"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/tencent/weixin/service" //注意导入Thrift生成的接口包
	"weixinsdk/src/utils"
)

func (this *WxServiceThrift) UnifiedOrder(orderParam *z_weixin_service.UnifiedOrderParam) (*z_weixin_service.UnifiedOrderResponse, error) {

	if !utils.ParamCheckForUnifiedOrder(orderParam) {
		return nil, errors.New("请检查订单必传参数")
	}

	//基础数据
	m_AppID := zconfig.CFG.MustValue("service", "AppID", "")
	m_MchID := zconfig.CFG.MustValue("service", "MchID", "")
	m_MchKey := zconfig.CFG.MustValue("service", "MchKey", "")
	nonceStr := utils.GenerateNonceString()
	if orderParam.Appid == "" {
		orderParam.Appid = m_AppID //设置APPID
	}
	orderParam.NonceStr = nonceStr //设置随机字符串
	orderParam.MchID = m_MchID     //设置商户ID

	//订单结构体转MAP
	orderMap := utils.Struct2Map(orderParam)

	//签名数据
	sign, err := utils.GenerateSignString(orderMap, m_MchKey)

	if err != nil {
		return nil, err
	}
	orderParam.Sign = sign //设置签名

	//生成XML
	requestXml, err := utils.GenerateRequestXml(orderParam)

	if err != nil {
		return nil, err
	}

	fmt.Println(requestXml)
	// //发起请求
	// r, err := http.Post(ApiUrlMap["UnifiedOrder"], "text/xml", strings.NewReader(requestXml))

	// if err != nil {
	// 	return nil, err
	// }

	// defer r.Body.Close()
	// body, err := ioutil.ReadAll(r.Body)

	// if err != nil {
	// 	return nil, err
	// }

	// var res WeResOrder

	// xml.Unmarshal([]byte(string(body)), &res)

	// if res.ReturnCode != "SUCCESS" || res.ResultCode != "SUCCESS" {
	// 	return nil, errors.New(res.ReturnMsg)
	// }
	logger.MyLogger.Info(requestXml)
	return nil, nil
}

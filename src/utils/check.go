package utils

import (
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/tencent/weixin/service" //注意导入Thrift生成的接口包
)

//订单参数检查
func ParamCheckForUnifiedOrder(orderParam *z_weixin_service.UnifiedOrderParam) bool {
	if orderParam.Body == "" {
		return false
	}
	if orderParam.OutTradeNo == "" {
		return false
	}
	if orderParam.TotalFee == "" {
		return false
	}
	if orderParam.SpbillCreateIp == "" {
		return false
	}
	if orderParam.NotifyUrl == "" {
		return false
	}
	if orderParam.TradeType == "" {
		return false
	}

	return true
}

package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"sort"
	"strings"
	zconfig "weixinsdk/src/config"
	"weixinsdk/src/lib"
	"weixinsdk/src/logger"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/tencent/weixin/service" //注意导入Thrift生成的接口包
	"weixinsdk/src/utils"

	"github.com/zouhuigang/package/zhttp"
)

//因为thrift生成的不带xml格式，所以在这里转换一下
//订单参数
type UnifiedOrderParam struct {
	Appid          string `xml:"appid"`            // 微信支付分配的公众账号ID（企业号corpid即为此appId）
	MchId          string `xml:"mch_id"`           // 微信支付分配的商户号
	DeviceInfo     string `xml:"device_info"`      // 自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"
	NonceStr       string `xml:"nonce_str"`        // 随机字符串，长度要求在32位以内 - https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_3
	Sign           string `xml:"sign"`             // 通过签名算法计算得出的签名值 - https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_3
	SignType       string `xml:"sign_type"`        // 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
	Body           string `xml:"body"`             // 商品简单描述，该字段请按照规范传递 - https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_2
	Detail         string `xml:"detail"`           // 商品详细描述，对于使用单品优惠的商户，改字段必须按照规范上传 - https://pay.weixin.qq.com/wiki/doc/api/danpin.php?chapter=9_102&index=2
	Attach         string `xml:"attach"`           // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用
	OutTradeNo     string `xml:"out_trade_no"`     // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|* 且在同一个商户号下唯一 - https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_2
	FeeType        string `xml:"fee_type"`         // 符合ISO 4217标准的三位字母代码，默认人民币：CNY - https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_2
	TotalFee       string `xml:"total_fee"`        // 订单总金额，单位为分 - https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_2
	SpbillCreateIp string `xml:"spbill_create_ip"` // 支持IPV4和IPV6两种格式的IP地址。调用微信支付API的机器IP
	TimeStart      string `xml:"time_start"`       // 订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010 - https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_2
	TimeExpired    string `xml:"time_expired"`     // 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。订单失效时间是针对订单号而言的，由于在请求支付的时候有一个必传参数prepay_id只有两小时的有效期，所以在重入时间超过2小时的时候需要重新请求下单接口获取新的prepay_id - https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_2
	GoodsTag       string `xml:"goods_tag"`        // 订单优惠标记，使用代金券或立减优惠功能时需要的参数 - https://pay.weixin.qq.com/wiki/doc/api/tools/sp_coupon.php?chapter=12_1
	NotifyUrl      string `xml:"notify_url"`       // 异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数
	TradeType      string `xml:"trade_type"`       // JSAPI -JSAPI支付,NATIVE -Native支付,APP -APP支付 - https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_2
	ProductId      string `xml:"product_id"`       // trade_type=NATIVE时，此参数必传。此参数为二维码中包含的商品ID，商户自行定义
	LimitPay       string `xml:"limit_pay"`        // 上传此参数no_credit--可限制用户不能使用信用卡支付
	Openid         string `xml:"openid"`           // trade_type=JSAPI时（即JSAPI支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识 - https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_4
	Receipt        string `xml:"receipt"`          // Y，传入Y时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效
	SceneInfo      string `xml:"scene_info"`       // 该字段常用于线下活动时的场景信息上报，支持上报实际门店信息，商户也可以按需求自己上报相关信息。该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }}
}

type UnifiedOrderResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Appid      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	TradeType  string `xml:"trade_type"`
	PrepayId   string `xml:"prepay_id"`
	CodeUrl    string `xml:"code_url"`
}

type WXPayNotify struct {
	ReturnCode    string `xml:"return_code"`
	ReturnMsg     string `xml:"return_msg"`
	Appid         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	DeviceInfo    string `xml:"device_info"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	ResultCode    string `xml:"result_code"`
	ErrCode       string `xml:"err_code"`
	ErrCodeDes    string `xml:"err_code_des"`
	Openid        string `xml:"openid"`
	IsSubscribe   string `xml:"is_subscribe"`
	TradeType     string `xml:"trade_type"`
	BankType      string `xml:"bank_type"`
	TotalFee      int64  `xml:"total_fee"`
	FeeType       string `xml:"fee_type"`
	CashFee       int64  `xml:"cash_fee"`
	CashFeeType   string `xml:"cash_fee_type"`
	CouponFee     int64  `xml:"coupon_fee"`
	CouponCount   int64  `xml:"coupon_count"`
	CouponID0     string `xml:"coupon_id_0"`
	CouponFee0    int64  `xml:"coupon_fee_0"`
	TransactionID string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	Attach        string `xml:"attach"`
	TimeEnd       string `xml:"time_end"`
}

func (this *WxServiceThrift) UnifiedOrder(orderParam *z_weixin_service.UnifiedOrderParam) (*z_weixin_service.UnifiedOrderResponse, error) {
	lg := fmt.Sprintf(" UnifiedOrder Receiving:%v", orderParam)
	logger.MyLogger.Debug(lg)

	if !lib.ParamCheckForUnifiedOrder(orderParam) {
		return nil, errors.New("请检查订单必传参数")
	}

	my_orderParam := new(UnifiedOrderParam)
	err := utils.StructCopy(orderParam, my_orderParam)
	if err != nil {
		return nil, errors.New("结构体复制失败")
	}

	//基础数据
	m_AppID := zconfig.CFG.MustValue("service", "AppID", "")
	m_MchID := zconfig.CFG.MustValue("service", "MchID", "")
	m_MchKey := zconfig.CFG.MustValue("service", "MchKey", "")
	nonceStr := utils.GenerateNonceString()
	if my_orderParam.Appid == "" {
		my_orderParam.Appid = m_AppID //设置APPID
	}
	my_orderParam.NonceStr = nonceStr //设置随机字符串
	my_orderParam.MchId = m_MchID     //设置商户ID

	//订单结构体转MAP
	orderMap := utils.Struct2Map(my_orderParam)

	//签名数据
	sign, err := utils.GenerateSignString(orderMap, m_MchKey)

	if err != nil {
		return nil, err
	}
	my_orderParam.Sign = sign //设置签名

	//记录发给微信服务器的数据
	lg = fmt.Sprintf("UnifiedOrder post xml data to weixin:%s", my_orderParam)
	logger.MyLogger.Info(lg)

	//发起请求
	responseXml := new(UnifiedOrderResponse)
	requrl := fmt.Sprintf("%s", zconfig.SERVICE_APIURL_PAY_UNIFIEDORDER)
	err = zhttp.PostXmlWithUnmarshal(requrl, my_orderParam, responseXml) //因为thrift没用生成xml支持的结构体

	lg = fmt.Sprintf("UnifiedOrder weixin response:%s", responseXml)
	logger.MyLogger.Debug(lg)
	// if res.ReturnCode != "SUCCESS" || res.ResultCode != "SUCCESS" {
	// 	return nil, errors.New(res.ReturnMsg)
	// }

	response := z_weixin_service.NewUnifiedOrderResponse()
	err = utils.StructCopy(responseXml, response)
	if err != nil {
		return nil, errors.New("结构体复制失败")
	}

	return response, err
}

//支付回调验证签名是否一致
//https://blog.csdn.net/xyzhaopeng/article/details/50386349
//https://blog.csdn.net/TauCrus/article/details/90241918
func (this *WxServiceThrift) WxpayParseAndVerifySign(xmlBytes []byte) (*z_weixin_service.WXPayNotify, error) {

	//1.解析消息

	var wxn WXPayNotify
	err := xml.Unmarshal(xmlBytes, &wxn)
	if err != nil {
		return nil, err
	}

	//2.验证签名数据,微信支付验证签名是否正确，防止别人任意回调
	//订单结构体转MAP
	var reqMap map[string]interface{}
	reqMap = make(map[string]interface{}, 0)
	reqMap["return_code"] = wxn.ReturnCode
	reqMap["return_msg"] = wxn.ReturnMsg
	reqMap["appid"] = wxn.Appid
	reqMap["mch_id"] = wxn.MchID
	reqMap["nonce_str"] = wxn.NonceStr
	reqMap["result_code"] = wxn.ResultCode
	reqMap["openid"] = wxn.Openid
	reqMap["is_subscribe"] = wxn.IsSubscribe
	reqMap["trade_type"] = wxn.TradeType
	reqMap["bank_type"] = wxn.BankType
	reqMap["total_fee"] = wxn.TotalFee
	reqMap["fee_type"] = wxn.FeeType
	reqMap["cash_fee"] = wxn.CashFee
	reqMap["cash_fee_type"] = wxn.CashFeeType
	reqMap["transaction_id"] = wxn.TransactionID
	reqMap["out_trade_no"] = wxn.OutTradeNo
	reqMap["attach"] = wxn.Attach
	reqMap["time_end"] = wxn.TimeEnd
	err = wxpayVerifySign(reqMap, wxn.Sign)
	if err != nil {
		return nil, err
	}

	//订单结构体转MAP
	response := z_weixin_service.NewWXPayNotify()
	err = utils.StructCopy(wxn, response)
	if err != nil {
		return nil, errors.New("结构体复制失败")
	}
	return response, nil
}

//微信支付计算签名的函数
func wxpayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}

	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}

func wxpayVerifySign(needVerifyM map[string]interface{}, sign string) error {
	m_MchKey := zconfig.CFG.MustValue("service", "MchKey", "")

	signCalc := wxpayCalcSign(needVerifyM, m_MchKey)

	// fmt.Println("========", signCalc, sign)
	if sign == signCalc {
		return nil
	}

	return errors.New("签名校验失败" + signCalc)
}

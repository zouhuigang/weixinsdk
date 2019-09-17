namespace go tencent.weixin.service
namespace php tencent.weixin.service
namespace py tencent.weixin.service

//微信统一下单参数
//https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
struct UnifiedOrderParam{
    1:string Appid
    2:string MchId
    3:string DeviceInfo
    4:string NonceStr
    5:string Sign
    6:string SignType
    7:string Body
    8:string Detail
    9:string Attach
    10:string OutTradeNo
    11:string FeeType
    12:string TotalFee
    13:string SpbillCreateIp
    14:string TimeStart
    15:string TimeExpire
    16:string GoodsTag
    17:string NotifyUrl
    18:string TradeType
    19:string ProductId
    20:string LimitPay
    21:string Openid
    22:string Receipt
    23:string Scene_info
}


//下单成功/失败返回
struct UnifiedOrderResponse {
	1:string ReturnCode
	2:string ReturnMsg
    //以下字段在return_code为SUCCESS的时候有返回
    3:string Appid
    4:string MchId
    5:string DeviceInfo
    6:string NonceStr
    7:string Sign
    8:string ResultCode
    9:string ErrCode
    10:string ErrCodeDes
    //以下字段在return_code 和result_code都为SUCCESS的时候有返回
    11:string TradeType
    12:string PrepayId
    13:string CodeUrl
}




struct JsApiParameters {
    //微信jsapiPay,因为package为关键字，所以用prepay_id代替
    1:string prepayId
    2:string nonceStr
    3:i64 timeStamp
    4:string sign
    5:string appId
    6:string paySign

}

//微信支付回调消息
/*
 <xml><appid><![CDATA[wxadd472a86212b893]]></appid>
<bank_type><![CDATA[LQT]]></bank_type>
<cash_fee><![CDATA[1]]></cash_fee>
<fee_type><![CDATA[CNY]]></fee_type>
<is_subscribe><![CDATA[Y]]></is_subscribe>
<mch_id><![CDATA[1555438221]]></mch_id>
<nonce_str><![CDATA[HFN4TCYHK6M7TJVVX7KBNNRC4H8QR6LA]]></nonce_str>
<openid><![CDATA[ozc4fs5UHdrxc7hVRoC42Yv2qB1k]]></openid>
<out_trade_no><![CDATA[WE0190916180864191]]></out_trade_no>
<result_code><![CDATA[SUCCESS]]></result_code>
<return_code><![CDATA[SUCCESS]]></return_code>
<sign><![CDATA[AC49F7C8102BC9EB007CCE2AC1A5F61F]]></sign>
<time_end><![CDATA[20190917135711]]></time_end>
<total_fee>1</total_fee>
<trade_type><![CDATA[JSAPI]]></trade_type>
<transaction_id><![CDATA[4200000409201909179443368122]]></transaction_id>
</xml>
*/
struct WXPayNotify{
	1:string return_code
	2:string return_msg
	3:string appid
	4:string mch_id
	5:string device_info
	6:string nonce_str
    7:string sign
	8:string result_code
    9:string err_code
    10:string err_code_des
    11:string openid
	12:string is_subscribe
	13:string trade_type
    14:string bank_type
    15:i64 total_fee
    16:string fee_type
	17:i64 cash_fee
    18:string cash_fee_type
    19:i64 coupon_fee
    20:i64 coupon_count
    21:string coupon_id_0
    22:i64 coupon_fee_0
    23:string transaction_id
    24:string out_trade_no
    25:string attach
    26:string time_end
}
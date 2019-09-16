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
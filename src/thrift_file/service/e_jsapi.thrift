namespace go tencent.weixin.service
namespace php tencent.weixin.service
namespace py tencent.weixin.service


       
/*
jsapi 微信支付
公众号id	appId	是	String(16)	wx8888888888888888	商户注册具有支付权限的公众号成功后即可获得
时间戳	timeStamp	是	String(32)	1414561699	当前的时间，其他详见时间戳规则
随机字符串	nonceStr	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
订单详情扩展字符串	package	是	String(128)	prepay_id=123456789	统一下单接口返回的prepay_id参数值，提交格式如：prepay_id=***
签名方式	signType	是	String(32)	MD5	签名类型，默认为MD5，支持HMAC-SHA256和MD5。注意此处需与统一下单的签名类型一致
签名	paySign	是	String(64)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
*/
struct JsapiSignData{
    1: string jsapi_ticket,
    2: string noncestr,
    3:i64 timestamp,
    4: string url,
    5: string sign,
    6: string appid,
    //微信jsapiPay,因为package为关键字，所以用prepay_id代替
    7:string prepay_id
    8:string paySign
    9:string signType
}

//首先定义一个UnifyOrderReq用于填入我们要传入的参数。
// type UnifyOrderReq struct {
//     Appid            string `xml:"appid"`            //公众账号ID
//     Body             string `xml:"body"`             //商品描述
//     Mch_id           string `xml:"mch_id"`           //商户号
//     Nonce_str        string `xml:"nonce_str"`        //随机字符串
//     Notify_url       string `xml:"notify_url"`       //通知地址
//     Trade_type       string `xml:"trade_type"`       //交易类型
//     Spbill_create_ip string `xml:"spbill_create_ip"` //支付提交用户端ip
//     Total_fee        int    `xml:"total_fee"`        //总金额
//     Out_trade_no     string `xml:"out_trade_no"`     //商户订单号
//     Sign             string `xml:"sign"`             //签名
//     Openid           string `xml:"openid"`           //购买商品的用户wxid
// }
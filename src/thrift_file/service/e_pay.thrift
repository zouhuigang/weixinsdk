namespace go tencent.weixin.service
namespace php tencent.weixin.service
namespace py tencent.weixin.service

//微信统一下单参数 
//https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
struct UnifiedOrderParam{
    1:string appid
    2:string mch_id
    3:string device_info
    4:string nonce_str
    5:string sign
    6:string sign_type
    7:string body
    8:string detail
    9:string attach
    10:string out_trade_no
    11:string fee_type
    12:string total_fee
    13:string spbill_create_ip
    14:string time_start
    15:string time_expire
    16:string goods_tag
    17:string notify_url
    18:string trade_type
    19:string product_id
    20:string limit_pay
    21:string openid
    22:string receipt
    23:string scene_info
}


//下单成功/失败返回
struct UnifiedOrderResponse {
	1:string return_code
	2:string return_msg
    //以下字段在return_code为SUCCESS的时候有返回
    3:string appid
    4:string mch_id
    5:string device_info
    6:string nonce_str
    7:string sign
    8:string result_code
    9:string err_code
    10:string err_code_des
    //以下字段在return_code 和result_code都为SUCCESS的时候有返回
    11:string trade_type
    12:string prepay_id
    13:string code_url
}
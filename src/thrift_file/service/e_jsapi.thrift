namespace go tencent.weixin.service
namespace php tencent.weixin.service
namespace py tencent.weixin.service

struct JsapiSignData{
    1: string jsapi_ticket,
    2: string noncestr,
    3:i64 timestamp,
    4: string url,
    5: string sign,
    6: string appid,
}

namespace go tencent.weixin.service
namespace php tencent.weixin.service
namespace py tencent.weixin.service

struct QrRespone{
    1:string Url//二维码链接
    2:string Content//二维码内容
    3:string ReqContent//请求的内容
    4:i32 ExpireSeconds//二维码有效期
}
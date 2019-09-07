namespace go tencent.weixin.service
namespace php tencent.weixin.service
namespace py tencent.weixin.service

struct WxResponse{
    1:i32 Errcode
	2:string Errmsg
	3:i64 MsgID
}
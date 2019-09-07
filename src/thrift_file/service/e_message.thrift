namespace go tencent.weixin.service
namespace php tencent.weixin.service
namespace py tencent.weixin.service

struct IsWeixinServerData{
    1:bool is_server
    2:string echostr
}
// 微信服务器推送过来的消息(事件)的合集.
struct MixedMessage  {
	//通用
	1:string ToUserName
	2:string FromUserName
	3:i64 CreateTime
	4:string MsgType
	5:i64 MsgId   // 消息id，64位整型
	6:i64 AgentId //跳转链接时所在的企业应用ID

	//文本消息
	7:optional string Content  // 文本消息内容
	// 图片消息
	8:optional string PicURL
	9:optional string MediaId

	// 音频消息
	10:optional string Format

	// 视频或短视频
	11:optional string ThumbMediaId

	// 地理位置消息
	12:double Location_X
	13:double Location_Y
	14:i32 Scale
	15:string Label

	// 链接消息
	16:string Title
	17:string Description
	18:string Url

	// 事件
	19:string Event
}


struct ParseTemplateToMixedMessagesData{
    1:required  string responeMessageType //消息类型
    2:required  MixedMessage responeMessage //消息内容
}


//发送模板消息https://blog.csdn.net/junmoxi/article/details/85471285
struct KeyWordData {
	1:string value
	2:string color
}
struct TemplateData  {
	1:KeyWordData  first
	2:KeyWordData  keyword1
	3:KeyWordData  keyword2
	4:KeyWordData  keyword3
	5:KeyWordData  keyword4
	6:KeyWordData  keyword5
	7:KeyWordData  remark
}

struct TemplateMsgData  {
	1:string touser           //接收者的OpenID
	2:string template_id   //模板消息ID
	3:string FormID 
	4:string url        //点击后跳转链接
	5:MiniprogramData miniprogram //点击跳转小程序
	6:TemplateData data
}
struct MiniprogramData  {
	1:string appid
	2:string pagepath
}

//发送回复模板消息
struct SendTemplateResponseData {
	1:i32 Errcode
	2:string Errmsg
	3:i64 MsgID
}
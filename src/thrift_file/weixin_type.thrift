namespace go weixin.service
namespace php weixin.service
namespace py weixin.service



/**
map(t,t): 键类型为t，值类型为t的kv对，键不容许重复。对应c++中的map, Java的HashMap, PHP 对应 array, Python/Ruby 的dictionary
 http://www.cpper.cn/2016/03/18/develop/Thrift-The-Missing-Guide/
 http://ju.outofmemory.cn/entry/263563
bool 布尔型
byte ８位整数
i16  16位整数
i32  32位整数
i64  64位整数
double 双精度浮点数
string 字符串
binary 字节数组
list<i16> List集合，必须指明泛型
map<string, string> Map类型，必须指明泛型
set<i32> Set集合，必须指明泛型 
 */
struct Article{
 1: i32 id, 
 2: string title,
 3: string content,
 4: string author,
}

struct JsapiSignData{
    1: string jsapi_ticket,
    2: string noncestr,
    3:i64 timestamp,
    4: string url,
    5: string sign,
    6: string appid,
}

struct IsWeixinServerData{
    1:bool is_server
    2:string echostr
}

//typedef map<string, string> Data

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


//用户信息
struct UserInfo{
	1:string openid
	2:string nickname
	3:i32 sex
	4:string province
	5:string city
	6:string country
	7:string headimgurl
	8:list<string> privilege
	9:string unionid
}

//返回认证数据
struct AuthCodeURLData{
	1:string url
	2:string state
}
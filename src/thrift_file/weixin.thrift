namespace go weixin.service
namespace php weixin.service

/**
 * 结构体定义
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

const map<string,string> MAPCONSTANT = {'hello':'world', 'goodnight':'moon'}

//微信服务号接口
service WxServiceThrift{        
        list<string> CallBack(1:i64 callTime, 2:string name, 3:map<string, string> paramMap),
        void put(1: Article newArticle),
        //获取全局access_token
        string GetAccessToken(),
        //获取jsapi_ticket
        string GetJsapiTicket(),
        JsapiSignData JsapiSign(1:string url)
        //是微信服务器过来的请求，用来验证消息回调
        IsWeixinServerData IsWeixinServer(1:string token, 2:string echostr, 3:string signature, 4:string  timestamp, 5:string  nonce)
}
namespace go weixin.service
namespace php weixin.service
namespace py weixin.service
include "weixin_type.thrift"

//微信服务号接口
service WxServiceThrift{        
        list<string> CallBack(1:i64 callTime, 2:string name, 3:map<string, string> paramMap),
        void put(1: weixin_type.Article newArticle),
        //获取全局access_token
        string GetAccessToken(),
        //获取jsapi_ticket
        string GetJsapiTicket(),
        weixin_type.JsapiSignData JsapiSign(1:string url),
        //是微信服务器过来的请求，用来验证消息回调
        weixin_type.IsWeixinServerData IsWeixinServer(1:string token, 2:string echostr, 3:string signature, 4:string  timestamp, 5:string  nonce),
        //解析模板消息
        weixin_type.ParseTemplateToMixedMessagesData  ParseTemplateToMixedMessages(1:string body),
        //发送模板消息
        weixin_type.SendTemplateResponseData SendTmplateMessage(1:weixin_type.TemplateMsgData tpl),
        //拼接模板消息
        string GetTextXml(1:string fromUserName, 2:string toUserName, 3:string content),
        //构造网页授权地址
        weixin_type.AuthCodeURLData AuthCodeURL(1:string redirectURL, 2:string scope),
        //根据code和snsapi_base获取用户信息
        weixin_type.UserInfo GetUserInfoBySnsapiBase(1:string code),
        //根据code和snsapi_userinfo获取用户信息
        weixin_type.UserInfo GetUserInfoBySnsapiUserinfo(1:string code),
        //根据openid拉取用户信息
        weixin_type.UserInfo GetUserInfoByOpenid(1:string openid),
       


}
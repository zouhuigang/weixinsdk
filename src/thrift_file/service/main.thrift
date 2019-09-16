namespace go tencent.weixin.service
namespace php tencent.weixin.service
namespace py tencent.weixin.service
include "e_curstom.thrift"
include "e_jsapi.thrift"
include "e_menu.thrift"
include "e_message.thrift"
include "e_oauth.thrift"
include "e_userinfo.thrift"
include "e_respone.thrift"
include "e_pay.thrift"

//微信服务号接口
service WxServiceThrift{        
        list<string> CallBack(1:i64 callTime, 2:string name, 3:map<string, string> paramMap),
        void put(1: e_curstom.Article newArticle),
        //获取全局access_token
        string GetAccessToken(),
        //获取jsapi_ticket
        string GetJsapiTicket(),
        e_jsapi.JsapiSignData JsapiSign(1:string url),
        //是微信服务器过来的请求，用来验证消息回调
        e_message.IsWeixinServerData IsWeixinServer(1:string token, 2:string echostr, 3:string signature, 4:string  timestamp, 5:string  nonce),
        //解析模板消息
        e_message.ParseTemplateToMixedMessagesData  ParseTemplateToMixedMessages(1:string body),
        //发送模板消息
        e_message.SendTemplateResponseData SendTmplateMessage(1:e_message.TemplateMsgData tpl),
        //拼接模板消息
        string GetTextXml(1:string fromUserName, 2:string toUserName, 3:string content),
        //转发给客服
        string TransferCustomerService(1:string fromUserName, 2:string toUserName, 3:string kfAccount),
        //构造网页授权地址
        e_oauth.AuthCodeURLData AuthCodeURL(1:string redirectURL, 2:string scope),
        //根据code和snsapi_base获取用户信息
        e_userinfo.UserInfo GetUserInfoBySnsapiBase(1:string code),
        //根据code和snsapi_userinfo获取用户信息
        e_userinfo.UserInfo GetUserInfoBySnsapiUserinfo(1:string code),
        //根据openid拉取用户信息
        e_userinfo.UserInfo GetUserInfoByOpenid(1:string openid),
        //创建菜单
        e_respone.WxResponse CreateMenu(1:e_menu.menu menu),
        e_respone.WxResponse CreateMenuByJson(1:binary  menuJsonBytes),
        //统一下单
        e_pay.UnifiedOrderResponse UnifiedOrder(1:e_pay.UnifiedOrderParam orderParam), 
        //返回js api pay参数
        e_pay.JsApiParameters GetJsApiParameters(1:e_pay.UnifiedOrderResponse unifiedOrderResult),


}
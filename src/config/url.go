package config

const (
	//全局access_token
	SERVICE_APIURL_ACCESS_TOKEN = `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential`
	//jsapi_ticket
	SERVICE_APIURL_JSAPI_TICKET = `https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi`
	//模板消息
	SERVICE_APIURL_SEND_TEMPLATE = `https://api.weixin.qq.com/cgi-bin/message/template/send`
)

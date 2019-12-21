package config

const (
	//全局access_token
	SERVICE_APIURL_ACCESS_TOKEN = `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential`
	//jsapi_ticket
	SERVICE_APIURL_JSAPI_TICKET = `https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi`
	//模板消息
	SERVICE_APIURL_SEND_TEMPLATE = `https://api.weixin.qq.com/cgi-bin/message/template/send`
	//网页授权获取access_token
	SERVICE_APIURL_OAUTH_ACCESS_TOKEN = `https://api.weixin.qq.com/sns/oauth2/access_token?grant_type=authorization_code`
	//获取用户信息
	SERVICE_APIURL_USER_INFO = `https://api.weixin.qq.com/cgi-bin/user/info`
	//网页授权获取用户信息，仅仅在snsapi_userinfo可用
	SERVICE_APIURL_USER_INFO_SNSAPI_USERINFO = `https://api.weixin.qq.com/sns/userinfo`
	//创建菜单
	SERVICE_APIURL_MENU_CREATE = `https://api.weixin.qq.com/cgi-bin/menu/create`
	//统一下单
	SERVICE_APIURL_PAY_UNIFIEDORDER = `https://api.mch.weixin.qq.com/pay/unifiedorder`
	//二维码生成,会返回一个ticket
	SERVICE_APIURL_QRCODE_CREATE = `https://api.weixin.qq.com/cgi-bin/qrcode/create`
	//二维码显示,根据ticket展示的是一张图片
	SERVICE_APIURL_QRCODE_SHOW = `https://mp.weixin.qq.com/cgi-bin/showqrcode`
	//获取素材列表
	SERVICE_APIURL_BATCHGET_MATERIAL = `https://api.weixin.qq.com/cgi-bin/material/batchget_material`
	//获取素材总数
	SERVICE_APIURL_MATERIAL_COUNT = `https://api.weixin.qq.com/cgi-bin/material/get_materialcount`
	//上传素材
	SERVICE_APIURL_MATERIAL_ADD = `https://api.weixin.qq.com/cgi-bin/material/add_material`
	//客服发消息
	SERVICE_APIURL_KEFU_SEND = `https://api.weixin.qq.com/cgi-bin/message/custom/send`
)

package structure

//错误的报错 https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1433747234
//https://godoc.org/github.com/mitchellh/mapstructure,不加mapstructure:",squash"访问不了
type ErrorBaseInfo struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

//全局也叫普通的access_token
type AccessToken struct {
	ErrorBaseInfo `mapstructure:",squash"`
	Expires_in    int64  `json:"expires_in"`
	Access_token  string `json:"access_token"`
	NowTimeStamp  int64  `json:"NowTimeStamp"`
}

//jsapi_ticket
type JsapiTicket struct {
	ErrorBaseInfo `mapstructure:",squash"`
	Ticket        string `json:"ticket"`
	Expires_in    int64  `json:"expires_in"`
	NowTimeStamp  int64  `json:"NowTimeStamp"`
}

//网页授权access_token,每个用户管理维护一个网页授权access_token
type OauthAccessToken struct {
	ErrorBaseInfo `mapstructure:",squash"`
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
	Openid        string `json:"openid"`
	Scope         string `json:"scope"`
	Expires_in    int64  `json:"expires_in"`
	NowTimeStamp  int64  `json:"NowTimeStamp"`
}

package structure

//全局也叫普通的access_token
type AccessToken struct {
	Expires_in   int64  `json:"expires_in"`
	Access_token string `json:"access_token"`
	NowTimeStamp int64  `json:"NowTimeStamp"`
}

//jsapi_ticket
type JsapiTicket struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Ticket       string `json:"ticket"`
	Expires_in   int64  `json:"expires_in"`
	NowTimeStamp int64  `json:"NowTimeStamp"`
}

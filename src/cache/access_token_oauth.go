package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	zconfig "weixinsdk/src/config"
	"weixinsdk/src/structure"

	"github.com/zouhuigang/mapstructure"
	"github.com/zouhuigang/package/zhttp"
	"github.com/zouhuigang/package/ztime"
)

/*
微信网页授权(OAuth2.0 )获得的access_token
有效期:7200 暂时不缓存
https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140842
https://github.com/bmob/goLib/blob/0447b3e9cf79b92f72e7c355d23f2c29d07555be/chanxuehong/wechat/mp/user/oauth2/oauth2.go
*/

//全局结构体
var (
	MyOauthAccessToken structure.OauthAccessToken
)

//存储中的key值
const m_OAUTH_ACCESS_TOKEN_KEY = `weixin_service_oauth_access_token`

//得到oauth access_token
func GetOauthAccessToken(code string) (structure.OauthAccessToken, error) {

	MyOauthAccessToken, err = initOauthAccessToken(code)

	return MyOauthAccessToken, err
}

func initOauthAccessToken(code string) (structure.OauthAccessToken, error) {

	token := structure.OauthAccessToken{}

	var m_AppID string = zconfig.CFG.MustValue("service", "AppID", "")
	var m_AppSecret string = zconfig.CFG.MustValue("service", "AppSecret", "")

	requrl := fmt.Sprintf("%s&appid=%s&secret=%s&code=%s", zconfig.SERVICE_APIURL_OAUTH_ACCESS_TOKEN, m_AppID, m_AppSecret, url.QueryEscape(code))
	json_str := zhttp.HttpGet(requrl)
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(json_str), &m); err != nil {
		return token, err
	}

	if err := mapstructure.Decode(m, &token); err != nil {
		return token, err
	}

	if token.Errcode != 0 {
		wxMsg := fmt.Sprintf("Errcode:%d,Errmsg:%s", token.Errcode, token.Errmsg)
		return token, errors.New(wxMsg)
	}

	//存储
	token.NowTimeStamp = ztime.NowTimeStamp()
	fmt.Printf("weixin server get oauth_access_token: %s \n", json_str)

	return token, nil

}

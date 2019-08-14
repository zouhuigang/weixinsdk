package service

import (
	"fmt"
	"net/url"
	zcache "weixinsdk/src/cache"
	zconfig "weixinsdk/src/config"
	"weixinsdk/src/logger"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/weixin/service" //注意导入Thrift生成的接口包
	zutils "weixinsdk/src/utils"

	"github.com/zouhuigang/package/zhttp"
)

// 构造请求用户授权获取code的地址.
//  appId:       公众号的唯一标识
//  redirectURL: 授权后重定向的回调链接地址
//               如果用户同意授权，页面将跳转至 redirect_uri/?code=CODE&state=STATE。
//               若用户禁止授权，则重定向后不会带上code参数，仅会带上state参数redirect_uri?state=STATE
//  scope:       应用授权作用域，
//               snsapi_base （不弹出授权页面，直接跳转，只能获取用户openid），
//               snsapi_userinfo （弹出授权页面，可通过openid拿到昵称、性别、所在地。
//               并且，即使在未关注的情况下，只要用户授权，也能获取其信息）
//  state:       重定向后会带上state参数，开发者可以填写a-zA-Z0-9的参数值，最多128字节
func (this *WxServiceThrift) AuthCodeURL(redirectURL, scope string) (*z_weixin_service.AuthCodeURLData, error) {
	var m_AppID string = zconfig.CFG.MustValue("service", "AppID", "")
	var state string = string(zutils.NewToken())

	data := z_weixin_service.NewAuthCodeURLData()
	data.URL = "https://open.weixin.qq.com/connect/oauth2/authorize" +
		"?appid=" + url.QueryEscape(m_AppID) +
		"&redirect_uri=" + url.QueryEscape(redirectURL) +
		"&response_type=code&scope=" + url.QueryEscape(scope) +
		"&state=" + url.QueryEscape(state) +
		"#wechat_redirect"
	data.State = state
	return data, nil
}

//根据openid拉取用户信息
func (this *WxServiceThrift) GetUserInfoByOpenid(openid string) (userinfo *z_weixin_service.UserInfo, err error) {

	access_token, err := zcache.GetAccessToken()
	if err != nil {
		return nil, err
	}
	requrl := fmt.Sprintf("%s?access_token=%s&openid=%s&lang=zh_CN", zconfig.SERVICE_APIURL_USER_INFO, access_token, openid)

	err = zhttp.GETWithUnmarshal(requrl, &userinfo)

	lg := fmt.Sprintf("GetUserInfoByOpenid get %v", userinfo)
	logger.MyLogger.Error(lg)

	return
}

//无感授权
func (this *WxServiceThrift) GetUserInfoBySnsapiBase(code string) (userinfo *z_weixin_service.UserInfo, err error) {
	//根据code获取openid
	oauth, err := zcache.GetOauthAccessToken(code)
	if err != nil {
		return nil, err
	}

	//拉取用户信息
	userinfo, err = this.GetUserInfoByOpenid(oauth.Openid)
	return
}

//用户授权
func (this *WxServiceThrift) GetUserInfoBySnsapiUserinfo(code string) (userinfo *z_weixin_service.UserInfo, err error) {
	//根据code获取openid
	oauth, err := zcache.GetOauthAccessToken(code)
	if err != nil {
		return nil, err
	}

	//拉取用户信息
	requrl := fmt.Sprintf("%s?access_token=%s&openid=%s&lang=zh_CN", zconfig.SERVICE_APIURL_USER_INFO_SNSAPI_USERINFO, oauth.Access_token, oauth.Openid)

	err = zhttp.GETWithUnmarshal(requrl, &userinfo)

	lg := fmt.Sprintf("GetUserInfoBySnsapiUserinfo get %v", userinfo)
	logger.MyLogger.Error(lg)

	return userinfo, nil
}

package cache

/*
全局或普通的access_token，用来调用接口
有效期:7200

https请求方式: GET
https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET

返回的数据:
{
    "access_token": "24_Qjrq6HjWIs3H62J5QeQI-SaO9dEkinO-YkHFQ9g4-eya_wmPM7WEIMzuS1YO6jGeNfoIKP_afte88rwJZ0_YLnCSOsR7N7_jBJcsn56eLmCUi7yK2ZFUcUlEMvR7pbNp64OWqGNaz2P3y_KhMDIgAIAEVA",
    "expires_in": 7200
}

官方的注意事项:
1、建议公众号开发者使用中控服务器统一获取和刷新access_token，其他业务逻辑服务器所使用的access_token均来自于该中控服务器，不应该各自去刷新，否则容易造成冲突，导致access_token覆盖而影响业务；

2、目前access_token的有效期通过返回的expire_in来传达，目前是7200秒之内的值。中控服务器需要根据这个有效时间提前去刷新新access_token。在刷新过程中，中控服务器可对外继续输出的老access_token，此时公众平台后台会保证在5分钟内，新老access_token都可用，这保证了第三方业务的平滑过渡；

3、access_token的有效时间可能会在未来有调整，所以中控服务器不仅需要内部定时主动刷新，还需要提供被动刷新access_token的接口，这样便于业务服务器在API调用获知access_token已超时的情况下，可以触发access_token的刷新流程。
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	zconfig "weixinsdk/src/config"
	"weixinsdk/src/lib"
	"weixinsdk/src/logger"
	zstorage "weixinsdk/src/storage"

	"weixinsdk/src/structure"

	"github.com/zouhuigang/mapstructure"
	"github.com/zouhuigang/package/zhttp"
	"github.com/zouhuigang/package/zreg"
	"github.com/zouhuigang/package/ztime"
	// json数据解析包，其转化效率比官方自带的encoding/json包高
	// 建议使用该包进行json对象的处理
)

//全局结构体
var (
	MyAccessToken structure.AccessToken
	err           error
)

//存储中的key值
const m_ACCESS_TOKEN_KEY = `weixin_service_access_token`

//有效期，秒
const m_ACCESS_TOKEN_EXPIRES = 7200

//得到token
func GetAccessToken() (string, error) {
	logger.MyLogger.Debug("GetAccessToken 1")
	if !lib.CacheValid(MyAccessToken.Access_token, MyAccessToken.NowTimeStamp, MyAccessToken.Expires_in, 1200, "GetAccessToken") {
		logger.MyLogger.Debug("GetAccessToken 2")
		MyAccessToken, err = initAccessToken()
	}

	return MyAccessToken.Access_token, err
}

func initAccessToken() (structure.AccessToken, error) {
	logger.MyLogger.Debug("initAccessToken start")
	//读取storage中的数据
	token := structure.AccessToken{}

	//从缓存中读取值，如果存在，则还需要判断缓存是否有效
	err := lib.GetCacheFromStorageWithUnmarshal(m_ACCESS_TOKEN_KEY, &token)
	if err != nil {
		msg := fmt.Sprintf("initAccessToken GetCacheFromStorageWithUnmarshal  error: %s", err.Error())
		logger.MyLogger.Error(msg)
		return token, err
	}

	logger.MyLogger.Debug("initAccessToken2")

	if !lib.CacheValid(token.Access_token, token.NowTimeStamp, token.Expires_in, 1200, "initAccessToken") {
		//fmt.Println("====================")
		logger.MyLogger.Debug("initAccessToken3")
		var m_AppSecret string = zconfig.CFG.MustValue("service", "AppSecret", "")
		var m_AppID string = zconfig.CFG.MustValue("service", "AppID", "")
		requrl := fmt.Sprintf("%s&appid=%s&secret=%s", zconfig.SERVICE_APIURL_ACCESS_TOKEN, m_AppID, m_AppSecret)
		json_str := zhttp.HttpGet(requrl)
		m := make(map[string]interface{})
		if err := json.Unmarshal([]byte(json_str), &m); err != nil {
			//fmt.Println(err)
			msg := fmt.Sprintf("initAccessToken json Unmarshal error: %s", err.Error())
			logger.MyLogger.Error(msg)
			return token, err
		}

		if err := mapstructure.Decode(m, &token); err != nil {
			msg := fmt.Sprintf("initAccessToken Decode error: %s", err.Error())
			logger.MyLogger.Error(msg)
			return token, err
		}

		if token.Errcode != 0 {
			msg := fmt.Sprintf("initAccessToken token.Errcode:%d", token.Errcode)
			logger.MyLogger.Error(msg)

			wxMsg := fmt.Sprintf("Errcode:%d,Errmsg:%s", token.Errcode, token.Errmsg)
			return token, errors.New(wxMsg)
		}

		//存储
		token.NowTimeStamp = ztime.NowTimeStamp()
		tokenJson, err := json.Marshal(token)
		if err != nil {
			//fmt.Println("access_token json字符串错误")
			msg := fmt.Sprintf("initAccessToken json Marshal error: %s", err.Error())
			logger.MyLogger.Error(msg)
			return token, errors.New("access_token json字符串错误")
		}

		if !zreg.IsNull(token.Access_token) {
			err := zstorage.MyStorage.Set(m_ACCESS_TOKEN_KEY, string(tokenJson), m_ACCESS_TOKEN_EXPIRES)
			if err != nil {
				//fmt.Println("storage fail")
				msg := fmt.Sprintf("initAccessToken MyStorage.Set error: %s", err.Error())
				logger.MyLogger.Error(msg)

				return token, errors.New("storage fail")
			}

		}

		logger.MyLogger.Info("[initAccessToken] weixin server get access_token")

		return token, nil

	}

	logger.MyLogger.Info("[initAccessToken] storage get access_token")
	return token, nil

}

//被动刷新access_token
func getNewAccessToken() (string, error) {
	var err error
	MyAccessToken = structure.AccessToken{}
	MyAccessToken, err = initAccessToken()
	return MyAccessToken.Access_token, err
}

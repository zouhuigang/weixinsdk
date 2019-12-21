package cache

/*
生成签名之前必须先了解一下jsapi_ticket，jsapi_ticket是公众号用于调用微信JS接口的临时票据。
正常情况下，jsapi_ticket的有效期为7200秒，通过access_token来获取。
由于获取jsapi_ticket的api调用次数非常有限，频繁刷新jsapi_ticket会导致api调用受限，影响自身业务，开发者必须在自己的服务全局缓存jsapi_ticket 。
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	zconfig "weixinsdk/src/config"
	"weixinsdk/src/logger"
	zstorage "weixinsdk/src/storage"

	"weixinsdk/src/lib"
	"weixinsdk/src/structure"

	"github.com/zouhuigang/mapstructure"
	"github.com/zouhuigang/package/zhttp"
	"github.com/zouhuigang/package/zreg"
	"github.com/zouhuigang/package/ztime"
)

//全局结构体
var (
	MyJsapiTicket structure.JsapiTicket
)

//存储中的key值
const m_JSAPI_TICKET_KEY = `weixin_service_jsapi_ticket`

//有效期，秒
const m_JSAPI_TICKET_EXPIRES = 7200

//得到token
func GetJsapiTicket() (string, error) {

	if !lib.CacheValid(MyJsapiTicket.Ticket, MyJsapiTicket.NowTimeStamp, MyJsapiTicket.Expires_in, 1200, "GetJsapiTicket") {
		MyJsapiTicket, err = initJsapiTicket()
	}

	return MyJsapiTicket.Ticket, err
}

func initJsapiTicket() (structure.JsapiTicket, error) {

	//读取storage中的数据
	token := structure.JsapiTicket{}
	err := lib.GetCacheFromStorageWithUnmarshal(m_JSAPI_TICKET_KEY, &token)
	if err != nil {
		logger.MyLogger.Errorf(err.Error())
		return token, err
	}

	if !lib.CacheValid(token.Ticket, token.NowTimeStamp, token.Expires_in, 1200, "initJsapiTicket") {
		logger.MyLogger.Info("weixin server get jsapi_ticket \n")
		access_token, err := GetAccessToken()
		if err != nil {
			return token, err
		}
		requrl := fmt.Sprintf("%s&access_token=%s", zconfig.SERVICE_APIURL_JSAPI_TICKET, access_token)
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
		tokenJson, err := json.Marshal(token)
		if err != nil {
			return token, errors.New("jsapi_ticket json字符串错误")
		}
		fmt.Println(string(tokenJson))
		if !zreg.IsNull(token.Ticket) {
			err := zstorage.MyStorage.Set(m_JSAPI_TICKET_KEY, string(tokenJson), m_JSAPI_TICKET_EXPIRES)
			if err != nil {
				return token, errors.New("jsapi_ticket  storage fail")
			}

		}

		return token, nil
	}
	return token, nil

}

//被动刷新
func getNewJsapiTicket() (string, error) {
	MyJsapiTicket = structure.JsapiTicket{}
	MyJsapiTicket, err = initJsapiTicket()
	return MyJsapiTicket.Ticket, err
}

package utils

import (
	"errors"
	"fmt"
	"weixinsdk/src/logger"
	zstorage "weixinsdk/src/storage"

	"github.com/zouhuigang/package/zreg"
	"github.com/zouhuigang/package/ztime"
	// json数据解析包，其转化效率比官方自带的encoding/json包高
	// 建议使用该包进行json对象的处理
	jsoniter "github.com/json-iterator/go"
)

/*
判断缓存有效
myToken:获取的token字符串
myNowTimeStamp:获取token的时间
myExpiresIn:token有效期
advanceTime:提前xx秒刷新token,默认提前20分钟,1200,
source:来源的函数
*/
func CacheValid(myToken string, myNowTimeStamp int64, myExpiresIn int64, advanceTime int64, source string) (valid bool) {
	valid = true
	if advanceTime <= 60 {
		advanceTime = 1200
	}

	timestamp := ztime.NowTimeStamp()
	cache_timestamp := myNowTimeStamp + myExpiresIn - advanceTime

	if zreg.IsNull(myToken) { //如果Ticket为空，则重新获取
		valid = false
		logger.MyLogger.Debug("token is null,source:", source)
	} else if cache_timestamp <= timestamp { //如果到了有效期前20分钟，则重新获取
		valid = false
		logger.MyLogger.Debug("cache is not valid,source:", source)
	} else {
		lg := fmt.Sprintf("storage server,source:%s,cur_timestamp:%d,timestamp:%d", source, cache_timestamp, timestamp)
		logger.MyLogger.Debug(lg)
	}

	return
}

/*
从storage中取得值,写入到对应的结构体中
*/
func GetCacheFromStorageWithUnmarshal(key string, to interface{}) error {
	m_storage_json := zstorage.MyStorage.Get(key)
	if zreg.IsNull(m_storage_json) {
		return errors.New("get cache from stroage is null")
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.Unmarshal([]byte(m_storage_json), to); err != nil {
		return err
	}
	return nil
}

//缓冲时间
func GetAdvanceTime() {
	// 由于网络的延时, 分布式服务器之间的时间可能不是绝对同步, access_token 过期时间留了一个缓冲区;
	// switch {
	// case result.ExpiresIn > 60*60:
	// 	result.ExpiresIn -= 60 * 20
	// case result.ExpiresIn > 60*30:
	// 	result.ExpiresIn -= 60 * 10
	// case result.ExpiresIn > 60*15:
	// 	result.ExpiresIn -= 60 * 5
	// case result.ExpiresIn > 60*5:
	// 	result.ExpiresIn -= 60
	// case result.ExpiresIn > 60:
	// 	result.ExpiresIn -= 20
	// case result.ExpiresIn > 0:
	// default:
	// 	err = fmt.Errorf("invalid expires_in: %d", result.ExpiresIn)
	// 	return
	// }

}

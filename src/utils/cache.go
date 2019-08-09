package utils

import (
	"errors"
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

	if zreg.IsNull(myToken) { //如果Ticket为空，则重新获取
		valid = false
		logger.MyLogger.Info("token is null")
	} else if (myNowTimeStamp + myExpiresIn + advanceTime) >= ztime.NowTimeStamp() { //如果到了有效期前20分钟，则重新获取
		valid = false
		logger.MyLogger.Info("cache is not valid")
	}

	logger.MyLogger.Info("cache is  valid")
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
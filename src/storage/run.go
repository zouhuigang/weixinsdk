package storage

import (
	"errors"
	zconfig "weixinsdk/src/config"
)

//对外暴露的变量，需要在init中先初始化连接存储
var (
	MyStorage Exporter
)

//初始化
func Load() error {
	storage_type := zconfig.CFG.MustValue("parameter", "storage_type", "local")
	exporter := ExporterMap()[storage_type]

	var ok bool
	MyStorage, ok = exporter.(Exporter)
	if !ok {
		return errors.New("storage init fail")
	}
	err := MyStorage.New()
	if err != nil {
		return err
	}

	//fmt.Println("==========", storageInterface.Get("weixin_dev"), ok)
	return nil
}

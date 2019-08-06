package storage

/*
请在各个子类中实现接口对应的方法
*/
type Exporter interface {
	//新建和连接
	New() (err error)
	//添加
	Add() (err error)
	//修改
	Set(m_key string, m_value string, m_expiration int32) (err error)
	//删除
	Delete() (err error)
	//获取
	Get(key string) string
}

func ExporterMap() (m map[string]interface{}) {
	m = map[string]interface{}{
		"memcache": new(Zmemcache),
		"local":    new(Zmemcache),
		"redis":    new(Zredis),
	}
	return
}

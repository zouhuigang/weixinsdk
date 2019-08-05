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
	Set() (err error)
	//删除
	Delete() (err error)
	//获取
	Get(key string) string
}

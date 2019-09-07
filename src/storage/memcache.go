package storage

import (
	"errors"
	"fmt"
	zconfig "weixinsdk/src/config"
	"weixinsdk/src/logger"

	"github.com/bradfitz/gomemcache/memcache"
)

type Zmemcache struct {
	mc *memcache.Client
}

func (this *Zmemcache) New() error {
	ip := zconfig.CFG.MustValue("memcache", "server", "127.0.0.1")
	port := zconfig.CFG.MustValue("memcache", "port", "11211")
	server := fmt.Sprintf("%s:%s", ip, port)
	logger.MyLogger.Info(server)
	this.mc = memcache.New(server)
	if this.mc == nil {
		return errors.New("memcache New failed")
	}

	return nil
}

func (this *Zmemcache) Add() error {
	return nil
}

//默认0代表无有效期，实际上是30天
func (this *Zmemcache) Set(m_key string, m_value string, m_expiration int32) error {
	err := this.mc.Set(&memcache.Item{Key: m_key, Value: []byte(m_value), Expiration: m_expiration})
	if err != nil {
		msg := fmt.Sprintf("mem set %s  error: %s", m_key, err.Error())
		logger.MyLogger.Error(msg)
	}

	return err
}

func (this *Zmemcache) Delete() error {
	return nil
}

func (this *Zmemcache) Get(key string) string {
	it, err := this.mc.Get(key)
	if err != nil {
		msg := fmt.Sprintf("mem get %s data error: %s", key, err.Error())
		logger.MyLogger.Error(msg)
		return ""
	}
	// if string(it.Key) == key {
	// }
	return string(it.Value)
}

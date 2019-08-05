package storage

import (
	"errors"
	"fmt"
	zconfig "weixinsdk/src/config"

	"github.com/bradfitz/gomemcache/memcache"
)

type Zmemcache struct {
	mc *memcache.Client
}

func (this *Zmemcache) New() error {
	ip := zconfig.CFG.MustValue("memcache", "server", "127.0.0.1")
	port := zconfig.CFG.MustValue("memcache", "port", "11211")
	server := fmt.Sprintf("%s:%s", ip, port)
	this.mc = memcache.New(server)
	if this.mc == nil {
		return errors.New("memcache New failed")
	}

	return nil
}

func (this *Zmemcache) Add() error {
	return nil
}

func (this *Zmemcache) Set() error {
	this.mc.Set(&memcache.Item{Key: "weixin_dev", Value: []byte("my sdsasdsa")})
	return nil
}

func (this *Zmemcache) Delete() error {
	return nil
}

func (this *Zmemcache) Get(key string) string {
	it, err := this.mc.Get(key)
	if err != nil {
		return ""
	}
	// if string(it.Key) == key {
	// }
	return string(it.Value)
}

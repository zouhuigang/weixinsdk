package storage

//BoltDB
import (
	"errors"
	"time"
	zconfig "weixinsdk/src/config"

	"github.com/asdine/storm"
	bolt "go.etcd.io/bbolt"
)

const bucketName = "weixin"

type Zlocal struct {
	mc *storm.DB
}

func (this *Zlocal) conn() error {
	dbFile := zconfig.CFG.MustValue("local", "file", "weixin.db")
	var err error
	this.mc, err = storm.Open(dbFile, storm.BoltOptions(0600, &bolt.Options{Timeout: 1 * time.Second}))
	if err != nil {
		return errors.New("local db New failed")
	}
	return err
}

func (this *Zlocal) close() {
	this.mc.Close()
}

func (this *Zlocal) New() error {
	err := this.conn()
	defer this.close()
	return err
}

func (this *Zlocal) Add() error {
	return nil
}

func (this *Zlocal) Set(m_key string, m_value string, m_expiration int32) error {
	err := this.conn()
	defer this.close()
	if err != nil {
		return err
	}

	err = this.mc.Set(bucketName, m_key, m_value)
	return err
}

func (this *Zlocal) Delete() error {
	return nil
}

func (this *Zlocal) Get(key string) string {
	this.conn()
	defer this.close()

	// value, err := this.mc.GetBytes(bucketName, key)
	var value string
	err := this.mc.Get(bucketName, key, &value)
	if err != nil {
		return ""
	}
	//fmt.Println("===========", value)
	return value
}

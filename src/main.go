package main

import (
	"fmt"
	"log"
	"os"
	zconfig "weixinsdk/src/config"
	zstorage "weixinsdk/src/storage"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/tencent/weixin/service" //注意导入Thrift生成的接口包
	//zutils "weixinsdk/src/utils"
	"weixinsdk/src/core/service"
	"weixinsdk/src/logger"

	"git.apache.org/thrift.git/lib/go/thrift"
	//"github.com/sirupsen/logrus"
)

func init() {
	//加载全局配置
	err := zconfig.Load()
	if err != nil {
		log.Fatalf("加载配置文件失败:%s", err)
	}

	//加载日志
	err = logger.Load()
	if err != nil {
		log.Fatalf("加载配置文件失败:%s", err)
	}

	//初始化存储引擎
	err = zstorage.Load()
	if err != nil {
		log.Fatalf("存储初始化失败:%s", err)
	}

}

func main() {

	port := zconfig.CFG.MustInt("parameter", "rpc_port", 3333)

	ipser := fmt.Sprintf("0.0.0.0:%d", port)
	// 传输器,传输方式
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	//transportFactory := thrift.NewTBufferedTransportFactory(10000000)

	// 传输协议:二进制格式binary|json|simplejson|compact
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(ipser)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &service.WxServiceThrift{}
	processor := z_weixin_service.NewWxServiceThriftProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)

	env := zconfig.CFG.MustValue("parameter", "env", "")
	logger.MyLogger.Info("weixin core start success,server:", ipser, ",env:", env)
	server.Serve()

}

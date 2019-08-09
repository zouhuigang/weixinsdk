package main

import (
	"fmt"
	"log"
	"os"
	zconfig "weixinsdk/src/config"
	zstorage "weixinsdk/src/storage"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/weixin/service" //注意导入Thrift生成的接口包
	//zutils "weixinsdk/src/utils"
	"weixinsdk/src/core/service"
	"weixinsdk/src/logger"

	"git.apache.org/thrift.git/lib/go/thrift"
	//"github.com/sirupsen/logrus"
)

func init() {
	//加载全局配置
	err := zconfig.Load("/build/base.env.ini", "/build/dev.env.ini")
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

func main2() {
	// err := zstorage.MyStorage.Set("weixin_service_access_token", `{
	// 	"realname": "邹慧刚",
	// 	"mobile": 18117000088,
	// 	"sex": "F"
	// }`, 120)
	// if err != nil {
	// 	fmt.Println("storage fail", err.Error())
	// }

	// dat := zstorage.MyStorage.Get("weixin_service_access_token")
	// fmt.Printf("%s\n", dat)

	//=====
	// 24小时一个日志文件，最多存365天的日志文件。 再多了就删掉

	// for {
	// 	time.Sleep(time.Second * 3)
	// 	logger.MyLogger.Info("asdsadasdsad ")
	// 	logger.MyLogger.Errorf("=====================")
	// }
	logger.MyLogger.Info("success ")
}

func main() {

	//测试access_token
	//zcache.GetAccessToken()

	//测试jsapi_ticket
	//dat := zcache.GetJsapiTicket()
	//dat := zservice.Jsapi_sign(`http://c3.ab.51tywy.com/qrcode/test/test.html`)

	// port := zconfig.CFG.MustInt("parameter", "rpc_port", 3333)

	// is_success, msg, service := zutils.StartServerThrift(port)
	// if !is_success {
	// 	fmt.Printf("%s\n", msg)
	// 	os.Exit(3)
	// } else {
	// 	fmt.Printf("v%s启动成功[TCP4]，调用密码为:%s,端口为:%d\n", "1.0.0", "", port)
	// }

	logger.MyLogger.Info("info")
	logger.MyLogger.Error("error")
	logger.MyLogger.Warn("warn")
	logger.MyLogger.Debug("debug")
	logger.MyLogger.Trace("trace")
	//logger.MyLogger.Fatal("fatal")
	logger.MyLogger.Panic("panic")

	ipser := `0.0.0.0:3333`
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
	fmt.Printf("thrift server in %s\n", ipser)
	server.Serve()

}

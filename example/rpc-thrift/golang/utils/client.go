package utils

import (

	// "services/articles"
	// "services/comments"
	// "services/users"
	"fmt"
	"net"
	"os"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/weixin/service" //注意导入Thrift生成的接口包

	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	HOST = "127.0.0.1"
	PORT = "3333"
)

func GetWxServerClient() (*z_weixin_service.WxServiceThriftClient, *thrift.TSocket) {

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := z_weixin_service.NewWxServiceThriftClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
		os.Exit(1)
	}
	//defer transport.Close()

	return client, transport
}

// func GetArticleClient(host, port string, initialCap, maxCap int, timeout time.Duration) *articles.ArticleServiceClient {
// 	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
// 	client, err := NewThriftPoolClient(host, port, protocolFactory, protocolFactory, initialCap, maxCap)
// 	if err != nil {
// 		log.Panicln("GetArticleClient error: ", err)
// 	}
// 	client.SetTimeout(timeout)
// 	return articles.NewArticleServiceClient(client)
// }

// func GetCommentClient(host, port string, initialCap, maxCap int, timeout time.Duration) *comments.CommentServiceClient {
// 	protocolFactory := thrift.NewTCompactProtocolFactory()
// 	client, err := NewThriftPoolClient(host, port, protocolFactory, protocolFactory, initialCap, maxCap)
// 	if err != nil {
// 		log.Panicln("GetCommentClient error: ", err)
// 	}
// 	client.SetTimeout(timeout)
// 	return comments.NewCommentServiceClient(client)
// }

// func GetUserClient(host, port string, initialCap, maxCap int, timeout time.Duration) *users.UserServiceClient {
// 	protocolFactory := thrift.NewTCompactProtocolFactory()
// 	client, err := NewThriftPoolClient(host, port, protocolFactory, protocolFactory, initialCap, maxCap)
// 	if err != nil {
// 		log.Panicln("GetUserClient error: ", err)
// 	}
// 	client.SetTimeout(timeout)
// 	return users.NewUserServiceClient(client)
// }

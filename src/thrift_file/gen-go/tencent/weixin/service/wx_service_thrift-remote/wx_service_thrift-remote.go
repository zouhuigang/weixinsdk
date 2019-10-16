// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "tencent/weixin/service"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "   CallBack(i64 callTime, string name,  paramMap)")
  fmt.Fprintln(os.Stderr, "  void put(Article newArticle)")
  fmt.Fprintln(os.Stderr, "  string GetAccessToken()")
  fmt.Fprintln(os.Stderr, "  string GetJsapiTicket()")
  fmt.Fprintln(os.Stderr, "  JsapiSignData JsapiSign(string url)")
  fmt.Fprintln(os.Stderr, "  IsWeixinServerData IsWeixinServer(string token, string echostr, string signature, string timestamp, string nonce)")
  fmt.Fprintln(os.Stderr, "  ParseTemplateToMixedMessagesData ParseTemplateToMixedMessages(string body)")
  fmt.Fprintln(os.Stderr, "  SendTemplateResponseData SendTmplateMessage(TemplateMsgData tpl)")
  fmt.Fprintln(os.Stderr, "  string GetTextXml(string fromUserName, string toUserName, string content)")
  fmt.Fprintln(os.Stderr, "  string TransferCustomerService(string fromUserName, string toUserName, string kfAccount)")
  fmt.Fprintln(os.Stderr, "  AuthCodeURLData AuthCodeURL(string redirectURL, string scope)")
  fmt.Fprintln(os.Stderr, "  UserInfo GetUserInfoBySnsapiBase(string code)")
  fmt.Fprintln(os.Stderr, "  UserInfo GetUserInfoBySnsapiUserinfo(string code)")
  fmt.Fprintln(os.Stderr, "  UserInfo GetUserInfoByOpenid(string openid)")
  fmt.Fprintln(os.Stderr, "  WxResponse CreateMenu(menu menu)")
  fmt.Fprintln(os.Stderr, "  WxResponse CreateMenuByJson(string menuJsonBytes)")
  fmt.Fprintln(os.Stderr, "  UnifiedOrderResponse UnifiedOrder(UnifiedOrderParam orderParam)")
  fmt.Fprintln(os.Stderr, "  JsApiParameters GetJsApiParameters(UnifiedOrderResponse unifiedOrderResult)")
  fmt.Fprintln(os.Stderr, "  WXPayNotify WxpayParseAndVerifySign(string xmlBytes)")
  fmt.Fprintln(os.Stderr, "  string QrcodeShow(string qrJsonBytes)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := service.NewWxServiceThriftClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "CallBack":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "CallBack requires 3 args")
      flag.Usage()
    }
    argvalue0, err45 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err45 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    arg47 := flag.Arg(3)
    mbTrans48 := thrift.NewTMemoryBufferLen(len(arg47))
    defer mbTrans48.Close()
    _, err49 := mbTrans48.WriteString(arg47)
    if err49 != nil { 
      Usage()
      return
    }
    factory50 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt51 := factory50.GetProtocol(mbTrans48)
    containerStruct2 := service.NewWxServiceThriftCallBackArgs()
    err52 := containerStruct2.ReadField3(jsProt51)
    if err52 != nil {
      Usage()
      return
    }
    argvalue2 := containerStruct2.ParamMap
    value2 := argvalue2
    fmt.Print(client.CallBack(value0, value1, value2))
    fmt.Print("\n")
    break
  case "put":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Put requires 1 args")
      flag.Usage()
    }
    arg53 := flag.Arg(1)
    mbTrans54 := thrift.NewTMemoryBufferLen(len(arg53))
    defer mbTrans54.Close()
    _, err55 := mbTrans54.WriteString(arg53)
    if err55 != nil {
      Usage()
      return
    }
    factory56 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt57 := factory56.GetProtocol(mbTrans54)
    argvalue0 := service.NewArticle()
    err58 := argvalue0.Read(jsProt57)
    if err58 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Put(value0))
    fmt.Print("\n")
    break
  case "GetAccessToken":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetAccessToken requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetAccessToken())
    fmt.Print("\n")
    break
  case "GetJsapiTicket":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetJsapiTicket requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetJsapiTicket())
    fmt.Print("\n")
    break
  case "JsapiSign":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "JsapiSign requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.JsapiSign(value0))
    fmt.Print("\n")
    break
  case "IsWeixinServer":
    if flag.NArg() - 1 != 5 {
      fmt.Fprintln(os.Stderr, "IsWeixinServer requires 5 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    argvalue3 := flag.Arg(4)
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    fmt.Print(client.IsWeixinServer(value0, value1, value2, value3, value4))
    fmt.Print("\n")
    break
  case "ParseTemplateToMixedMessages":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ParseTemplateToMixedMessages requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.ParseTemplateToMixedMessages(value0))
    fmt.Print("\n")
    break
  case "SendTmplateMessage":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SendTmplateMessage requires 1 args")
      flag.Usage()
    }
    arg66 := flag.Arg(1)
    mbTrans67 := thrift.NewTMemoryBufferLen(len(arg66))
    defer mbTrans67.Close()
    _, err68 := mbTrans67.WriteString(arg66)
    if err68 != nil {
      Usage()
      return
    }
    factory69 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt70 := factory69.GetProtocol(mbTrans67)
    argvalue0 := service.NewTemplateMsgData()
    err71 := argvalue0.Read(jsProt70)
    if err71 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.SendTmplateMessage(value0))
    fmt.Print("\n")
    break
  case "GetTextXml":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "GetTextXml requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.GetTextXml(value0, value1, value2))
    fmt.Print("\n")
    break
  case "TransferCustomerService":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "TransferCustomerService requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.TransferCustomerService(value0, value1, value2))
    fmt.Print("\n")
    break
  case "AuthCodeURL":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "AuthCodeURL requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.AuthCodeURL(value0, value1))
    fmt.Print("\n")
    break
  case "GetUserInfoBySnsapiBase":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetUserInfoBySnsapiBase requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetUserInfoBySnsapiBase(value0))
    fmt.Print("\n")
    break
  case "GetUserInfoBySnsapiUserinfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetUserInfoBySnsapiUserinfo requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetUserInfoBySnsapiUserinfo(value0))
    fmt.Print("\n")
    break
  case "GetUserInfoByOpenid":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetUserInfoByOpenid requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetUserInfoByOpenid(value0))
    fmt.Print("\n")
    break
  case "CreateMenu":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CreateMenu requires 1 args")
      flag.Usage()
    }
    arg83 := flag.Arg(1)
    mbTrans84 := thrift.NewTMemoryBufferLen(len(arg83))
    defer mbTrans84.Close()
    _, err85 := mbTrans84.WriteString(arg83)
    if err85 != nil {
      Usage()
      return
    }
    factory86 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt87 := factory86.GetProtocol(mbTrans84)
    argvalue0 := service.NewMenu()
    err88 := argvalue0.Read(jsProt87)
    if err88 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CreateMenu(value0))
    fmt.Print("\n")
    break
  case "CreateMenuByJson":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CreateMenuByJson requires 1 args")
      flag.Usage()
    }
    argvalue0 := []byte(flag.Arg(1))
    value0 := argvalue0
    fmt.Print(client.CreateMenuByJson(value0))
    fmt.Print("\n")
    break
  case "UnifiedOrder":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "UnifiedOrder requires 1 args")
      flag.Usage()
    }
    arg90 := flag.Arg(1)
    mbTrans91 := thrift.NewTMemoryBufferLen(len(arg90))
    defer mbTrans91.Close()
    _, err92 := mbTrans91.WriteString(arg90)
    if err92 != nil {
      Usage()
      return
    }
    factory93 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt94 := factory93.GetProtocol(mbTrans91)
    argvalue0 := service.NewUnifiedOrderParam()
    err95 := argvalue0.Read(jsProt94)
    if err95 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.UnifiedOrder(value0))
    fmt.Print("\n")
    break
  case "GetJsApiParameters":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetJsApiParameters requires 1 args")
      flag.Usage()
    }
    arg96 := flag.Arg(1)
    mbTrans97 := thrift.NewTMemoryBufferLen(len(arg96))
    defer mbTrans97.Close()
    _, err98 := mbTrans97.WriteString(arg96)
    if err98 != nil {
      Usage()
      return
    }
    factory99 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt100 := factory99.GetProtocol(mbTrans97)
    argvalue0 := service.NewUnifiedOrderResponse()
    err101 := argvalue0.Read(jsProt100)
    if err101 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetJsApiParameters(value0))
    fmt.Print("\n")
    break
  case "WxpayParseAndVerifySign":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "WxpayParseAndVerifySign requires 1 args")
      flag.Usage()
    }
    argvalue0 := []byte(flag.Arg(1))
    value0 := argvalue0
    fmt.Print(client.WxpayParseAndVerifySign(value0))
    fmt.Print("\n")
    break
  case "QrcodeShow":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "QrcodeShow requires 1 args")
      flag.Usage()
    }
    argvalue0 := []byte(flag.Arg(1))
    value0 := argvalue0
    fmt.Print(client.QrcodeShow(value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
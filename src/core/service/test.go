package service

import (
	"fmt"
	"time"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/tencent/weixin/service" //注意导入Thrift生成的接口包
)

func (this *WxServiceThrift) CallBack(callTime int64, name string, paramMap map[string]string) (r []string, err error) {
	fmt.Println("-->from client Call:", time.Unix(callTime, 0).Format("2006-01-02 15:04:05"), name, paramMap)
	r = append(r, "key:"+paramMap["a"]+"    value:"+paramMap["b"])
	return
}

func (this *WxServiceThrift) Put(s *z_weixin_service.Article) (err error) {
	fmt.Printf("Article--->id: %d\tTitle:%s\tContent:%s\tAuthor:%s\n", s.ID, s.Title, s.Content, s.Author)
	return nil
}

package module

import (
	zcore "weixinsdk/src/core/service"
)

//链式调用，控制顺序，后期考虑用context控制
//https://blog.csdn.net/qq_34675369/article/details/101123353
type KeFu struct {
	handler      *zcore.WxServiceThrift
	SuccessCount int //执行成功的数量
}

type KeFuOption func() (*KeFu, error)

func (r *KeFu) WithCreate(mhandler *zcore.WxServiceThrift) KeFuOption {
	return func() (kf *KeFu, e error) {
		r.handler = mhandler
		return r, nil
	}
}

//发送文字消息
func (opt KeFuOption) WithSendText(content string) KeFuOption {
	return func() (request *KeFu, e error) {
		r, err := opt()
		if err != nil {
			return r, err
		}
		//发送
		_, err = r.handler.KefuSend([]byte(content))
		if err != nil {
			return r, err
		}
		r.SuccessCount += 1
		return r, err
	}
}

//发送图片消息
func (opt KeFuOption) WithSendImage(content string) KeFuOption {
	return func() (request *KeFu, e error) {
		r, err := opt()
		if err != nil {
			return r, err
		}
		//发送
		_, err = r.handler.KefuSend([]byte(content))
		if err != nil {
			return r, err
		}
		r.SuccessCount += 1
		return r, err
	}
}

func (opt KeFuOption) SuccessCount() (int, error) {
	r, err := opt()
	if err != nil {
		return 0, err
	}
	//fmt.Printf("执行成功数:%d", r.SuccessCount)
	return r.SuccessCount, nil
}

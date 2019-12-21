/*
全局返回状态码
https://github.com/grpc/grpc-go/blob/master/codes/codes.go
https://github.com/cli22/httpserver-test/blob/422e95e87619de79dd26b483438a831a884b5765/app/error/error.go
*/
package global

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type ErrCode uint32

const (
	ErrOk ErrCode = iota
	ErrUnknown
	ErrArgsInvalid
	ErrArgsEmpty
	ErrUserOrPasswdInvalid
	ErrUserNameEmpty
	ErrPasswdEmpty
	ErrSubmit
	ErrSubmitAjax
	ErrJwtUnauthorized
	ErrCaptchaErr
	ErrPasswdDiffer
	ErrEmailEmpty
	ErrEmail
	ErrEmailExist
	ErrRegisterFail
	ErrUserInfo
	ErrDate
	ErrYear
	ErrMonth
	ErrDay
	ErrTeacher
	ErrApi
	ErrSave
	ErrDel
	ErrSubject
	ErrType
	ErrSafeLock
)

//异常消息映射表
var ErrDescription = map[ErrCode]string{
	ErrOk:                  "成功",
	ErrUnknown:             "未知错误",
	ErrArgsInvalid:         "参数异常",
	ErrArgsEmpty:           "参数为空",
	ErrUserOrPasswdInvalid: "用户名或密码错误",
	ErrUserNameEmpty:       "用户名不能为空",
	ErrPasswdEmpty:         "密码不能为空",
	ErrSubmit:              "提交方式或数据格式有误",
	ErrSubmitAjax:          "非法的ajax请求",
	ErrJwtUnauthorized:     "无效的Token或Token已过期",
	ErrCaptchaErr:          "验证码错误，请重新获取",
	ErrPasswdDiffer:        "2次密码输入不一致",
	ErrEmailEmpty:          "邮箱不能为空",
	ErrEmail:               "邮箱格式错误",
	ErrEmailExist:          "邮箱已存在",
	ErrRegisterFail:        "注册失败,请稍后再试",
	ErrUserInfo:            "用户信息不存在",
	ErrDate:                "错误的日期",
	ErrYear:                "错误的年份",
	ErrMonth:               "错误的月份",
	ErrDay:                 "错误的天数",
	ErrTeacher:             "当前查询教师不存在或已离职",
	ErrApi:                 "接口不存在",
	ErrSave:                "保存失败，请稍后再试",
	ErrDel:                 "删除失败，请稍后再试",
	ErrSubject:             "错误的课程",
	ErrType:                "类型错误",
	ErrSafeLock:            "密码错误",
}

/*新的返回*/
func ResponeJson(ctx echo.Context, errCode ErrCode, data interface{}, args ...interface{}) error {
	var error_description string
	for _, arg := range args { //迭代不定参数
		switch arg.(type) {
		case string:
			error_description = arg.(string)
		default:
			error_description = ""
		}
	}

	if error_description == "" {
		error_description = ErrDescription[errCode]
	}

	result := map[string]interface{}{
		"error":             errCode,
		"error_description": error_description,
		"result":            data,
	}

	b, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return ctx.JSONBlob(http.StatusOK, b)
}

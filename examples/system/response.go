package system

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"

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
	ErrSql
	ErrUserInfo
	ErrImperfect
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
	ErrSql:                 "错误的SQL",
	ErrUserInfo:            "获取用户信息失败",
	ErrImperfect:           "未完善资料",
}

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

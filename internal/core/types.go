package core

import (
	"api-gateway/internal/configs"
)

type (
	App struct {
		Name         string
		Authenticate *Authenticate
		Perm         *Permission
	}

	Authenticate interface {
		Set(apps *configs.Apps)
		Add(appkey string, auth *configs.Auth)
		Verify(appkey, content, sign string) *Error
	}

	Permission interface {
		// Set 设置权限
		Set(permit *configs.Permit)
		// Verify 校验权限
		Verify(serviceName, method string) *Error
	}

	ServiceExecutor interface {
		Set(services *configs.Services)
		Call(serviceName, method string) (*Result, *Error)
	}

	Service interface {
		Set(service *configs.Service)
		Call(method string) (*Result, *Error)
	}

	Method interface {
		Call() (*Result, *Error)
	}
	// Request 请求
	Request struct {
		AppId     string `json:"appId"`
		Sign      string `json:"sign"`
		Timestamp int32  `json:"timestamp"`
		V         string `json:"version"`
		Req       struct {
			Service string `json:"service"`
			Method  string `json:"method"`
			Param   string `json:"param"`
		}
	}
	Error struct {
		Code    string `json:"code"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Sign    string `json:"sign,omitempty"`
	}
	// Result 返回值
	Result struct {
		//返回内容 当 status=true有效
		Data string `json:"data,omitempty"`
		//错误信息 当 status=false有效
		Message string `json:"message,omitempty"`
		//状态代码 当 status = false 有效
		Code string `json:"code"`
		//状态 true 表示 成功 false表示 失败
		Status bool   `json:"status"`
		Sign   string `json:"sign,omitempty"`
	}
)

func NewError(code string, message ...string) *Error {
	he := &Error{Code: code, Success: false}
	if len(message) > 0 {
		he.Message = message[0]
	}
	if len(message) > 1 {
		he.Sign = message[1]
	}
	return he
}

func NewResultErr(code string, message ...string) *Result {
	var he = &Result{Message: code, Status: false}
	if len(message) > 0 {
		he.Message = message[0]
	}
	if len(message) > 1 {
		he.Sign = message[1]
	}
	return he
}

func ErrResult(err *Error) *Result {
	return &Result{
		Message: err.Message,
		Status:  false,
		Code:    err.Code,
		Sign:    err.Sign,
	}
}

func NewResult(code, data, sign string) *Result {
	return &Result{Code: code, Status: true, Data: data, Sign: sign}
}

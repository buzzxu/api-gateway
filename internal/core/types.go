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
		Set(auth *configs.Auth)
		Verify(content, sign string) *Error
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
		AppId     string
		Sign      string
		Timestamp int32
		V         string
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
	}
	// Result 返回值
	Result struct {
		//返回内容 当 status=true有效
		Data string `json:"data"`
		//错误信息 当 status=false有效
		Message string `json:"message"`
		//状态代码 当 status = false 有效
		Code string `json:"code"`
		//状态 true 表示 成功 false表示 失败
		Status bool   `json:"status"`
		Sign   string `json:"sign"`
	}
)

func NewError(code string, message ...string) *Error {
	he := &Error{Code: code, Success: false}
	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}

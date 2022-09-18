package core

import (
	"api-gateway/internal/configs"
	"api-gateway/internal/core/auth"
	"github.com/buzzxu/ironman/logger"
)

var loader configs.Loader
var executor ServiceExecutor
var authenticate Authenticate

func init() {
	//加载配置
	loader = new(configs.YamlLoader)
	loader.Load()

	executor = new(DefaultServiceExecutor)
	executor.Set(loader.Services())

	authenticate = new(auth.AkSkAuthenticate)
	authenticate.Set(loader.Apps())
}

func Exec(request *Request) *Result {
	if request.AppId == "" {
		return NewResultErr("1001", "请设置appId")
	}
	logger.Of("request").Printf("appId: %s ,service: %s, method: %s", request.AppId, request.Req.Service, request.Req.Method)
	if request.Timestamp <= 0 {
		return NewResultErr("1001", "请设置timestamp")
	}
	if request.Sign == "" {
		return NewResultErr("1001", "请设置sign")
	}
	if request.V == "" {
		request.V = loader.Services().Version
	}
	appId := request.AppId
	if err := authenticate.Verify(appId, "", request.Sign); err != nil {
		return NewResultErr(err.Code, err.Message)
	}

	return nil
}

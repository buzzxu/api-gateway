package core

import (
	"api-gateway/internal/configs"
	"github.com/buzzxu/boys/common/ids"
	"github.com/buzzxu/ironman"
	"github.com/buzzxu/ironman/conf"
)

var loader configs.Loader
var executor ServiceExecutor
var idGen ids.SnowflakeID
var apps *Apps

func init() {
	//加载配置
	conf.LoadDefaultConf()
	ironman.RedisConnect()
	loader = new(configs.YamlLoader)
	loader.Load()

	apps = new(Apps)
	apps.Set(loader.Apps())

	executor = new(DefaultServiceExecutor)
	executor.Set(loader.Services())

	idGen = ids.SnowflakeID{}
	if err := idGen.Init(); err != nil {
		panic("IdGen init error.")
	}
}

func Exec(request *Request) *Result {
	requestId := idGen.Generate()
	req, app, result := request.DoIt(requestId)
	if result != nil {
		return result
	}
	//验证签名
	content := app.AppKey + request.Req + request.TimestampStr()

	if err := app.Verify(content, request.Sign); err != nil {
		return NewResultErr(requestId, err.Code, err.Message, app)
	}
	//req
	executor.Call(requestId, app, req.Service, req.Method, req.Param)
	return nil
}

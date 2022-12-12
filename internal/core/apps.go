package core

import (
	"api-gateway/internal/configs"
	"fmt"
	"github.com/buzzxu/boys/common/signature/aksk"
)

func (a *Apps) Set(apps *configs.Apps) {
	a.apps = make(map[string]*App)
	for _, _app := range apps.Apps {
		app := &App{
			Name:      _app.Name,
			AppKey:    _app.Auth.AppKey,
			AppSecret: _app.Auth.AppSecret,
		}
		app.Perm = &Permission{
			permit: _app.Permit,
		}
		a.apps[_app.Auth.AppKey] = app
	}
}

func (a *Apps) Get(appkey string) (*App, *Error) {
	if app, ok := a.apps[appkey]; ok {
		return app, nil
	}
	return nil, NewError("1006", fmt.Sprintf("appId: %s 不存在", appkey))
}

// Verify 验签
func (a *App) Verify(content, sign string) *Error {
	err := aksk.Verify(content, sign, a.AppSecret)
	if err != nil {
		return NewError("1007", err.Error())
	}
	return nil
}

func (a *App) Call(service, method string) (interface{}, *Error) {
	err := a.Perm.Verify(service, method)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

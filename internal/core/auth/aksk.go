package auth

import (
	"api-gateway/internal/configs"
	"api-gateway/internal/core"
	"fmt"
	"github.com/buzzxu/boys/common/signature/aksk"
)

type AkSkAuthenticate struct {
	apps map[string]*configs.Auth
}

func (s *AkSkAuthenticate) Set(apps *configs.Apps) {
	s.apps = make(map[string]*configs.Auth)
	for _, app := range apps.Apps {
		s.apps[app.Auth.AppKey] = app.Auth
	}
}
func (s *AkSkAuthenticate) Add(appkey string, auth *configs.Auth) {
	if s.apps == nil {
		s.apps = make(map[string]*configs.Auth)
	}
	s.apps[appkey] = auth
}

func (s *AkSkAuthenticate) Verify(appkey, content, sign string) *core.Error {
	if app, ok := s.apps[appkey]; ok {
		err := aksk.Verify(content, sign, app.AppSecret)
		if err != nil {
			return core.NewError("1007", err.Error())
		}
		return nil
	}
	return core.NewError("1006", fmt.Sprintf("appId: %s 不存在", appkey))
}

func (s AkSkAuthenticate) Get(appkey string) *configs.Auth {
	return s.apps[appkey]
}

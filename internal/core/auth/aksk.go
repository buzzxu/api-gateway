package auth

import (
	"api-gateway/internal/configs"
	"api-gateway/internal/core"
	"github.com/buzzxu/boys/common/signature/aksk"
)

type AkSkAuther struct {
	auth *configs.Auth
}

func (s *AkSkAuther) Set(auth *configs.Auth) {
	s.auth = auth
}

func (s *AkSkAuther) Verify(content, sign string) *core.Error {
	//timestamp
	//if
	err := aksk.Verify(content, sign, s.auth.AppSecret)
	if err != nil {
		return core.NewError("", err.Error())
	}
	return nil
}

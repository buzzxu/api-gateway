package core

import (
	"api-gateway/internal/configs"
	"net/http"
	"strconv"
)

type (
	DefaultServiceExecutor struct {
		services map[string]Service
	}
	DefaultService struct {
		auth    *configs.Auth
		methods map[string]Method
	}
)

func (s DefaultServiceExecutor) Set(services *configs.Services) {
	s.services = make(map[string]Service)
	for k, v := range services.Services {
		methods := make(map[string]Method)
		if v.Protocol == "http" {
			for _, _method := range v.Methods {
				httpMethod := http.MethodPost
				if _method.Method != "" {
					httpMethod = _method.Method
				}
				methods[_method.Method] = &HttpMethod{
					url:    v.Host + ":" + strconv.Itoa(v.Port) + _method.Url,
					method: httpMethod,
				}
			}
		}
		s.services[k] = &DefaultService{
			auth:    v.Auth,
			methods: methods,
		}
	}
}

// Call 调用服务类接口
func (s *DefaultServiceExecutor) Call(requestId string, app *App, serviceName, method, params string) *Result {
	//验证接口权限
	error := app.Perm.Verify(serviceName, method)
	if error != nil {
		return NewResultOfErr(requestId, error, app)
	}
	result, error := s.services[serviceName].Call(requestId, method, params)
	if error != nil {
		return NewResultOfErr(requestId, error, app)
	}
	return result
}

// Call 调用方法
func (service *DefaultService) Call(requestId, method, param string) (*Result, *Error) {
	return service.methods[method].Call(requestId, param)
}

func newMethod() *Method {
	return nil
}

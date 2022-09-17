package core

import "api-gateway/internal/configs"

type (
	DefaultServiceExecutor struct {
		services map[string]Service
	}
	DefaultService struct {
		auth    *configs.Auth
		methods map[string]*Method
	}
)

func (s DefaultServiceExecutor) Set(services *configs.Services) {
	s.services = make(map[string]Service)
	for k, v := range services.Services {
		methods := make(map[string]*Method)
		s.services[k] = &DefaultService{
			auth:    v.Auth,
			methods: methods,
		}
	}
}

func (s DefaultServiceExecutor) Call(serviceName, method string) (*Result, *Error) {
	return nil, nil
}

func (d DefaultService) Set(service *configs.Service) {
	//TODO implement me
	panic("implement me")
}

func (d DefaultService) Call(method string) (*Result, *Error) {
	//TODO implement me
	panic("implement me")
}

func newMethod() *Method {
	return nil
}

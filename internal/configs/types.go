package configs

import "time"

type (
	Loader interface {
		// Load 加载配置
		Load()
		// Apps 获取业务方数据
		Apps() *Apps
		// Services 获取服务类
		Services() *Services
		// Refresh 刷新数据
		Refresh(duration time.Duration)
	}

	Apps struct {
		Apps map[string]*App `yaml:"apps"`
	}

	App struct {
		Name    string  `yaml:"name"`
		Auth    *Auth   `yaml:"auth"`
		Expired string  `yaml:"expired"`
		Permit  *Permit `yaml:"permit"`
	}
	Permit struct {
		Scope    string `yaml:"scope,omitempty"`
		Services map[string]*struct {
			Allow []string `yaml:"allow"`
			Deny  []string `yaml:"deny"`
		}
	}

	Services struct {
		Version  string             `yaml:"v"`
		Services map[string]Service `yaml:"services"`
	}

	Service struct {
		Protocol    string            `yaml:"protocol"`
		Host        string            `yaml:"host"`
		Port        int               `yaml:"port"`
		Description string            `yaml:"description"`
		Auth        *Auth             `yaml:"auth"`
		Methods     map[string]Method `yaml:"methods"`
	}
	Auth struct {
		Type      string `yaml:"type"`
		Token     string `yaml:"token,omitempty"`
		AppKey    string `yaml:"appKey,omitempty"`
		AppSecret string `yaml:"appSecret,omitempty"`
	}

	Method struct {
		Url    string `yaml:"url"`
		Method string `yaml:"method,omitempty"`
		Limit  int    `yaml:"limit,omitempty"`
	}
)

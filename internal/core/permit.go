package core

import (
	"api-gateway/internal/configs"
	"fmt"
)

type DefPermission struct {
	permit *configs.Permit
}

// Set 设置权限
func (p DefPermission) Set(permit *configs.Permit) {
	p.permit = permit
}

// Verify 权限匹配
func (p DefPermission) Verify(serviceName, method string) *Error {
	if p.permit.Scope == "*" {
		//如果作用域是 *
		return nil
	}
	value, exists := p.permit.Services[serviceName]
	if exists {
		//存在 验证具体权限
		denySize := len(value.Deny)
		if denySize > 0 {
			//1.先验证deny 如果deny包含 * ，表示此service下的所有method 均被拒绝访问
			if denySize == 1 && value.Deny[0] == "*" {
				return NewError("1006", fmt.Sprintf("%s 被拒绝访问", serviceName))
			}
			//2.验证method是否被拒绝
			for _, val := range value.Deny {
				if val == method {
					return NewError("1006", fmt.Sprintf("service:  %s ,method: %s,被拒绝访问", serviceName, method))
				}
			}
		}
		allowSize := len(value.Allow)
		if allowSize > 0 {
			//3 验证 allow，如果allow包含*，表示此service下的所有method均被允许访问
			if allowSize == 1 && value.Allow[0] == "*" {
				return nil
			}
			//4 验证method是否同意
			for _, val := range value.Allow {
				if val == method {
					return nil
				}
			}
			//未包含就拒绝
			return NewError("1006", fmt.Sprintf("service:  %s ,method: %s,被拒绝访问", serviceName, method))
		}
	}
	//未配置某个service 允许客户端访问
	return nil
}

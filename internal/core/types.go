package core

import (
	"api-gateway/internal/configs"
	"context"
	"encoding/json"
	"fmt"
	"github.com/buzzxu/boys/common/signature/aksk"
	"github.com/buzzxu/boys/common/strs"
	"github.com/buzzxu/ironman"
	"github.com/buzzxu/ironman/logger"
	"strconv"
	"time"
)

type (
	Apps struct {
		apps map[string]*App
	}
	App struct {
		Name      string
		AppKey    string
		AppSecret string
		Perm      *Permission
		Executor  *ServiceExecutor
	}
	Authenticate interface {
		Set(apps *configs.Apps)
		Add(appkey string, auth *configs.Auth)
		Verify(appkey, content, sign string) *Error
		Get(appkey string) *configs.Auth
	}

	ServiceExecutor interface {
		Set(services *configs.Services)
		Call(requestId string, app *App, serviceName, method, params string) *Result
	}

	Service interface {
		Call(requestId, method, param string) (*Result, *Error)
	}
	Authenticater interface {
		GetToken() string
	}
	Method interface {
		Call(requestId, params string) (*Result, *Error)
	}
	// Request 请求
	Request struct {
		AppKey    string `json:"appKey"`
		Sign      string `json:"sign"`
		Timestamp int64  `json:"timestamp"`
		//AppId + Timestamp  10分钟之内不能重复
		Nonce string `json:"nonce"`
		V     string `json:"version"`
		Req   string `json:"req"`
	}
	Req struct {
		Service string `json:"service"`
		Method  string `json:"method"`
		Param   string `json:"param"`
	}
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	// Result 返回值
	Result struct {
		RequestId string `json:"requestId"`
		AppKey    string `json:"appKey"`
		//返回内容 当 status=true有效
		Data string `json:"data,omitempty"`
		//错误信息 当 status=false有效
		Message string `json:"message,omitempty"`
		//状态代码 当 status = false 有效
		Code string `json:"code"`
		//状态 true 表示 成功 false表示 失败
		Status bool   `json:"status"`
		Sign   string `json:"sign,omitempty"`
	}
	AppInfo struct {
		AppKey    string `yaml:"appKey,omitempty"`
		AppSecret string `yaml:"appSecret,omitempty"`
	}
)

func NewError(code, message string) *Error {
	return &Error{Code: code, Message: message}
}

func NewResultErr(requestId, code, message string, app *App) *Result {
	return NewResultOfErr(requestId, &Error{
		Code:    code,
		Message: message,
	}, app)
}
func NewResultErrOfNilSign(requestId, code, message string) *Result {
	return &Result{RequestId: requestId, Code: code, Status: false, Message: message}
}
func NewResultOfErr(requestId string, err *Error, app *App) *Result {
	var he = &Result{AppKey: app.AppKey, RequestId: requestId, Code: err.Code, Status: false, Message: err.Message}
	data, _ := json.Marshal(err)
	he.Data = string(data)
	sign, error := Sign(he.Data, app)
	if error != nil {
		return NewResultErr(requestId, "1008", fmt.Sprintf("签名失败,原因: %s", error.Error()), app)
	}
	he.Sign = sign
	return he
}

func ErrResult(err *Error) *Result {
	data, _ := json.Marshal(err)
	return &Result{
		Message: err.Message,
		Status:  false,
		Code:    err.Code,
		Data:    string(data),
	}
}

func New(requestId string, data interface{}, app *App) *Result {
	value, err := json.Marshal(data)
	if err != nil {
		return NewResultErr(requestId, "1011", fmt.Sprintf("返回数据序列化失败,原因: %s", err.Error()), app)
	}
	result := string(value)
	sign, err := Sign(result, app)
	if err != nil {
		return NewResultErr(requestId, "1008", fmt.Sprintf("签名失败,原因: %s", err.Error()), app)
	}
	return NewResult(requestId, app.AppKey, result, sign)
}
func NewResult(requestId, appKey, data, sign string) *Result {
	return &Result{RequestId: requestId, AppKey: appKey, Status: true, Data: data, Sign: sign}
}

func Sign(params string, app *App) (string, error) {
	sign, err := aksk.Signature(app.AppKey+params+strconv.FormatInt(time.Now().Unix(), 10), app.AppSecret)
	if err != nil {
		return "", err
	}
	return sign, nil
}

// To 转换Req参数
func (request *Request) To() (*Req, error) {
	req := &Req{}
	err := json.Unmarshal([]byte(request.Req), req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (request *Request) TimestampStr() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// CheckArgument 参数校验
func (request *Request) DoIt(requestId string) (*Req, *App, *Result) {
	if request.AppKey == "" {
		return nil, nil, NewResultErrOfNilSign(requestId, "1001", "请设置appKey")
	}
	if request.Nonce == "" {
		return nil, nil, NewResultErrOfNilSign(requestId, "1001", "请设置nonce")
	}
	if request.Timestamp <= 0 {
		return nil, nil, NewResultErrOfNilSign(requestId, "1001", "请设置timestamp")
	}
	if request.Sign == "" {
		return nil, nil, NewResultErrOfNilSign(requestId, "1001", "请设置sign")
	}
	if request.V == "" {
		//默认版本号
		request.V = loader.Services().Version
	}
	if request.Req == "" {
		return nil, nil, NewResultErrOfNilSign(requestId, "1001", "请设置req")
	}
	//AppId + Timestamp + nonce 5分钟之内不能重复
	reqKey := strconv.Itoa(int(strs.Hash32(strs.Concat(request.AppKey, request.TimestampStr(), request.Nonce))))
	if ironman.Redis.Exists(context.Background(), reqKey).Val() > 0 {
		return nil, nil, NewResultErrOfNilSign(requestId, "1012", "重复请求")
	}
	ironman.Redis.SetEX(context.Background(), reqKey, "0", 5*time.Minute)

	//获取app认证信息
	app, err := apps.Get(request.AppKey)
	if err != nil {
		return nil, nil, NewResultErr(requestId, err.Code, err.Message, app)
	}
	req, error := request.To()
	if err != nil {
		return nil, nil, NewResultErrOfNilSign(requestId, "1010", fmt.Sprintf("参数解析失败,原因: %s", error.Error()))
	}
	logger.Of("request").Printf("requestId: %s, appKey: %s ,service: %s, method: %s", requestId, request.AppKey, req.Service, req.Method)
	//验证Timestamp 时间间隔不能超过10分钟
	now := time.Now()
	timestamp := time.Unix(request.Timestamp, 0)
	if now.Before(timestamp) {
		return nil, nil, NewResultErr(requestId, "1009", "请求时间错误,请校验本地时间", app)
	}
	interval := now.Sub(timestamp)
	if interval.Minutes() > 10 {
		return nil, nil, NewResultErr(requestId, "1009", "请求时间已过期,拒绝响应", app)
	}
	return req, app, nil
}

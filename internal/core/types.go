package core

type Request struct {
	AppId     string
	Sign      string
	Timestamp int32
	V         string
	Req       struct {
		Service string `json:"service"`
		Method  string `json:"method"`
		Param   string `json:"param"`
	}
}

// Result 返回值
type Result struct {
	//返回内容 当 status=true有效
	Data string `json:"data"`
	//错误信息 当 status=false有效
	Message string `json:"message"`
	//状态代码 当 status = false 有效
	Code string `json:"code"`
	//状态 true 表示 成功 false表示 失败
	Status bool   `json:"status"`
	Sign   string `json:"sign"`
}

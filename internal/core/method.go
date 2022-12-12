package core

import (
	"encoding/json"
)

type HttpMethod struct {
	url    string
	method string
}

type ProtobufMethod struct {
}

func (m *HttpMethod) Call(requestId, params string) (*Result, *Error) {
	var jsonUnknow interface{}
	err := json.Unmarshal([]byte(params), &jsonUnknow)
	if err != nil {
		return nil, NewError("1010", err.Error())
	}
	//_, err := http.NewRequestWithContext(context.Background(), m.method, m.url, body)

	return nil, nil
}

func (m ProtobufMethod) Call(requestId, params string) *Result {
	return nil
}

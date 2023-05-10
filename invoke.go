package jsonlib

import (
	"context"
	"errors"
	"net/http"
)

type Invoke struct {
	Host           string              // rcp请求的地址(默认空)
	Path           string              // rpc请求的路径(默认空)
	Query          map[string][]string // rpc请求的参数(默认空)
	Metadata       map[string][]string // rpc请求的头部(默认空)
	RequestContent any                 // rpc请求的内容(默认空)
	ResponseStatus int                 // rpc响应状态(默认0不判断状态)
	ResponseResult any                 // rpc响应结果(默认空不解析结果)
}

type InvokeSetting struct {
	Trace func(req *http.Request, rsp *http.Response, err error) // 追踪打印请求函数
}

var ErrInvalidStatus = errors.New("invalid status")

type InvokeClient interface {
	Do(ctx context.Context, method string, invoke *Invoke) (int, error)
	GET(ctx context.Context, invoke *Invoke) (int, error)
	POST(ctx context.Context, invoke *Invoke) (int, error)
	PUT(ctx context.Context, invoke *Invoke) (int, error)
	DELETE(ctx context.Context, invoke *Invoke) (int, error)
}

func NewInvokeClien(c *InvokeSetting) InvokeClient {
	return nil
}

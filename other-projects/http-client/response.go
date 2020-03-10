package httpclient

import (
	"fmt"
	"net/http"
)

// ResponseCarrier 接口响应接口
type ResponseCarrier interface {
	GetErrorCode() int
	GetErrorMessage() string
}

// Response 接口响应通用字段
type Response struct {
	ErrorCode    int    `json:"errcode,omitempty"`
	ErrorMessage string `json:"errmsg,omitempty"`
}

// GetErrorCode 获取错误码
func (r *Response) GetErrorCode() int {
	return r.ErrorCode
}

// GetErrorMessage 获取错误信息
func (r *Response) GetErrorMessage() string {
	return r.ErrorMessage
}

// ResponseError 接口响应错误
type ResponseError struct {
	Code     int
	Message  string
	Response *http.Response
}

// Error 实现 error 接口
func (e *ResponseError) Error() string {
	resp := e.Response
	req := resp.Request
	return fmt.Sprintf("%s %s : [%d] %s", req.Method, req.URL.Path, e.Code, e.Message)
}

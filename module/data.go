package module

import "net/http"

// Data 数据接口类型
type Data interface {
	// Valid 判断数据是否有效
	Valid() bool
}

// Request 数据请求类型
type Request struct {
	// http请求
	httpReq *http.Request
	// 请求深度
	depth uint32
}

// NewRequest 创建一个新的请求实例
func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq: httpReq, depth: depth}
}

// HTTPReq 获取HTTP请求
func (req *Request) HTTPReq() *http.Request {
	return req.httpReq
}

// Depth 获取请求的深度
func (req *Request) Depth() uint32 {
	return req.depth
}

// Valid 判断请求是否有效
func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

// Response 数据响应类型
type Response struct {
	httpRes *http.Response
	depth   uint32
}

// NewResponse 创建响应实例
func NewResponse(httpRes *http.Response, depth uint32) *Response {
	return &Response{httpRes: httpRes, depth: depth}
}

// HTTPRes 获取http响应
func (res *Response) HTTPRes() *http.Response {
	return res.httpRes
}

// Depth 获取响应的深度
func (res *Response) Depth() uint32 {
	return res.depth
}

// Valid 判断响应是否有效
func (res *Response) Valid() bool {
	return res.httpRes != nil && res.httpRes.Body != nil
}

// Item 条目的类型
type Item map[string]interface{}

// Valid 判断条目是否有效
func (item Item) Valid() bool {
	return item != nil
}

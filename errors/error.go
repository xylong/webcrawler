package errors

import (
	"bytes"
	"fmt"
	"strings"
)

// ErrorType 代表错误类型
type ErrorType string

// 错误类型常量
const (
	// 下载器错误
	ERROR_TYPE_DOWNLOADER ErrorType = "downloader error"
	// 分析器错误
	ERROR_TYPE_ANALYZER ErrorType = "analyzer error"
	// 条目处理管道错误
	ERROR_TYPE_PIPELINE ErrorType = "pipeline error"
	// 调度器错误
	ERROR_TYPE_SCHEDULER ErrorType = "scheduler error"
)

// CrawlerError 爬虫错误的类型接口
type CrawlerError interface {
	// 获取错误类型
	Type() ErrorType
	// 获取错误提示信息
	Error() string
}

// NewCrawlerError 创建一个新的爬虫错误值
func NewCrawlerError(errorType ErrorType, msg string) CrawlerError {
	return &myCrawlerError{
		errType: errorType,
		errMsg:  strings.TrimSpace(msg),
	}
}

// myCrawlerError 爬虫错误的实现类型
type myCrawlerError struct {
	// 错误的类型
	errType ErrorType
	// 错误的提示信息
	errMsg string
	// 完整的错误提示信息
	fullErrMsg string
}

func (m *myCrawlerError) Type() ErrorType {
	return m.errType
}

func (m *myCrawlerError) Error() string {
	if m.fullErrMsg == "" {
		m.genFullErrMsg()
	}
	return m.fullErrMsg
}

// GenFullErrMsg 生成错误提示信息，并给响应字段赋值
func (m *myCrawlerError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("crawler error: ")
	if m.errType != "" {
		buffer.WriteString(string(m.errType))
		buffer.WriteString(": ")
	}
	buffer.WriteString(m.errMsg)
	m.fullErrMsg = fmt.Sprintf("%s", buffer.String())
	return
}

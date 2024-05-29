package xerror

import "github.com/SliverFlow/ksmall/monserver/common/constant"

// XError 自定义错误
type XError struct {
	Code    uint
	Message string
}

// New 创建自定义错误
func New(code uint, message string) *XError {
	return &XError{
		Code:    code,
		Message: message,
	}
}

// Error 实现error接口
func (e *XError) Error() string {
	return e.Message
}

// NewWithCode 通过code创建自定义错误
func NewWithCode(code uint) *XError {
	message, ok := constant.MessageCodeMap[code]
	if !ok {
		message = "未知错误"
	}

	return &XError{
		Code:    code,
		Message: message,
	}
}

// NewWithMessage 通过message创建自定义错误
func NewWithMessage(message string) *XError {
	return &XError{
		Code:    constant.RequestFailedCode,
		Message: message,
	}
}

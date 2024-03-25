package zerror

import "github.com/SliverFlow/ksmall/common/constant"

// ZError 自定义错误
type ZError struct {
	Code    uint
	Message string
}

// New 创建自定义错误
func New(code uint, message string) *ZError {
	return &ZError{
		Code:    code,
		Message: message,
	}
}

// Error 实现error接口
func (e *ZError) Error() string {
	return e.Message
}

// NewWithCode 通过code创建自定义错误
func NewWithCode(code uint) *ZError {
	message, ok := constant.MessageCodeMap[code]
	if !ok {
		message = "未知错误"
	}

	return &ZError{
		Code:    code,
		Message: message,
	}
}

// NewWithMessage 通过message创建自定义错误
func NewWithMessage(message string) *ZError {
	return &ZError{
		Code:    constant.RequestFailedCode,
		Message: message,
	}
}

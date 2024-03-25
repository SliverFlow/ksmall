package response

import (
	"errors"
	"github.com/SliverFlow/ksmall/common/constant"
	"github.com/SliverFlow/ksmall/common/zerror"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SuccessResponse struct {
	Code    uint        `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ErrorResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

const SuccessCode = 0

// OkWithDetail 请求成功，返回数据和消息
func OkWithDetail(data interface{}, message string, c *gin.Context) {
	c.JSON(http.StatusOK, &SuccessResponse{Code: SuccessCode, Data: data, Message: message})
}

// Ok 请求成功，返回消息
func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, &SuccessResponse{Code: SuccessCode, Data: nil, Message: "success"})
}

// OkWithMessage 请求成功，返回消息
func OkWithMessage(message string, c *gin.Context) {
	c.JSON(http.StatusOK, &SuccessResponse{Code: SuccessCode, Data: nil, Message: message})
}

// OkWithData 请求成功，返回数据
func OkWithData(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, &SuccessResponse{Code: SuccessCode, Data: data, Message: "success"})
}

// Fail 请求失败，返回消息
func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, &ErrorResponse{Code: constant.RequestFailedCode, Message: "fail"})
}

// FailWithMessage 请求失败，返回消息
func FailWithMessage(message string, c *gin.Context) {
	c.JSON(http.StatusOK, &ErrorResponse{Code: constant.RequestFailedCode, Message: message})

}

// FailWithDetail 请求失败，返 code 和消息
func FailWithDetail(code uint, message string, c *gin.Context) {
	c.JSON(http.StatusOK, &ErrorResponse{Code: code, Message: message})
}

func FailWithError(err error, c *gin.Context) {
	var zerr *zerror.ZError
	if errors.As(err, &zerr) {
		// err is a ZError, return its message and code
		c.JSON(http.StatusOK, &ErrorResponse{Code: zerr.Code, Message: zerr.Message})
		return
	}
	// err is not a ZError, return a generic error message
	c.JSON(http.StatusOK, &ErrorResponse{Code: constant.RequestFailedCode, Message: "服务器开小差了，请稍后再试"})
}

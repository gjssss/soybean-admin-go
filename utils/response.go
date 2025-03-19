package utils

import "github.com/gin-gonic/gin"

type Response[T any] struct {
	Code string `json:"code" binding:"required"`
	Data T      `json:"data" binding:"required"`
	Msg  string `json:"msg"  binding:"required"`
}

func NewResponse[T any](code string, data T, msg string) Response[T] {
	return Response[T]{Code: code, Data: data, Msg: msg}
}

func NewSuccessResponse[T any](data T) Response[T] {
	return NewResponse("0000", data, "success")
}
func NewErrorResponse(msg string) Response[string] {
	return NewResponse("0000", "", msg)
}
func NewLogoutModelResponse(code string, msg string) Response[string] {
	return NewResponse("8888", "", msg)
}

// Success 直接向客户端返回成功响应
func Success[T any](c *gin.Context, data T) {
	c.JSON(200, NewSuccessResponse(data))
}

// Fail 直接向客户端返回错误响应
func Fail(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, NewErrorResponse(message))
}

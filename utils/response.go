package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS = http.StatusOK
	ERROR   = http.StatusBadRequest
)

var msg = map[int]string{
	SUCCESS: "success",
	ERROR:   "fail",
}

func GetMsg(code int) string {
	messages, ok := msg[code]
	if ok {
		return messages
	}
	return ""
}

func Result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(SUCCESS, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Success(c *gin.Context) {
	Result(SUCCESS, GetMsg(SUCCESS), map[string]interface{}{}, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, GetMsg(ERROR), map[string]interface{}{}, c)
}

func SuccessWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, message, map[string]interface{}{}, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, message, map[string]interface{}{}, c)
}

func SuccessWithData(data interface{}, message string, c *gin.Context) {
	if message != "" {
		Result(SUCCESS, message, data, c)
	} else {
		Result(SUCCESS, GetMsg(SUCCESS), data, c)
	}

}

func FailWithData(data interface{}, message string, c *gin.Context) {
	if message != "" {
		Result(ERROR, message, data, c)
	} else {
		Result(ERROR, GetMsg(ERROR), data, c)
	}
}

package handler

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func Respond(c *gin.Context, httpStatus, code int, msg string, data any) {
	c.JSON(httpStatus, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

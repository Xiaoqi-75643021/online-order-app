package handler

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Content any    `json:"content,omitempty"`
}

func Respond(c *gin.Context, httpStatus, code int, msg string, content any) {
	c.JSON(httpStatus, Response{
		Code:    code,
		Message: msg,
		Content: content,
	})
}

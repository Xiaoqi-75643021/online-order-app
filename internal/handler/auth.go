package handler

import (
	"net/http"
	"online-ordering-app/internal/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	type LoginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}

	token, err := service.Login(req.Username, req.Password)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "登录失败", gin.H{"error": err.Error()})
		return
	}
	Respond(c, http.StatusOK, 0, "登录成功", gin.H{"token": token})
}

func Register(c *gin.Context) {
	type RegisterRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role"`
	}
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}

	token, err := service.Register(req.Username, req.Password, req.Role)
	if err != nil {
		statusCode := http.StatusInternalServerError
		var code int
		switch err {
		case service.ErrUserAlreadyExists:
			statusCode = http.StatusBadRequest
			code = 3
		case service.ErrCreateUserFailed, service.ErrGenerateTokenFailed:
			statusCode = http.StatusInternalServerError
			code = 4
		}
		Respond(c, statusCode, code, "注册失败", gin.H{
			"field": "username",
			"error": err.Error(),
		})
		return
	}
	Respond(c, http.StatusOK, 0, "注册成功", gin.H{"token": token})
}

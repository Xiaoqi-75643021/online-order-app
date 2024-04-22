package handler

import (
	"net/http"
	"online-ordering-app/internal/service"

	"github.com/gin-gonic/gin"
)

func UpdateAdminName(c *gin.Context) {
	type request struct {
		NewUsername string `json:"newUsername" binding:"required"`
	}
	var req request
	if err := c.ShouldBind(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{
			"field": "username",
			"error": err.Error(),
		})
		return
	}
	userID, _ := c.Get("userID")
	err := service.UpdateAdminName(userID.(uint), req.NewUsername)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "用户名更新失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "用户名更新成功", nil)
}

func UpdateAdminPassword(c *gin.Context) {
	type request struct {
		NewPassword string `json:"newPassword" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{
			"field": "password",
			"error": err.Error(),
		})
		return
	}

	userID, _ := c.Get("userID")
	err := service.UpdateAdminPassword(userID.(uint), req.NewPassword)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "密码更新失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "密码更新成功", nil)
}


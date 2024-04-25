package handler

import (
	"net/http"
	"online-ordering-app/internal/service"

	"github.com/gin-gonic/gin"
)

func UpdateUsername(c *gin.Context) {
	type request struct {
		NewUsername string `json:"newUsername" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{
			"field": "newUsername",
			"error": err.Error(),
		})
		return
	}
	userID, _ := c.Get("user_id")
	err := service.UpdateUsername(userID.(uint), req.NewUsername)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "用户名更新失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "用户名更新成功", nil)
}

func UpdatePassword(c *gin.Context) {
	type request struct {
		NewPassword string `json:"newPassword" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{
			"field": "newPassword",
			"error": err.Error(),
		})
		return
	}

	userID, _ := c.Get("user_id")
	err := service.UpdatePassword(userID.(uint), req.NewPassword)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "密码更新失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "密码更新成功", nil)
}

func RechargeBalance(c *gin.Context) {

}

func DeductBalance(c *gin.Context) {

}

package handler

import (
	"net/http"
	"online-ordering-app/internal/service"

	"github.com/gin-gonic/gin"
)

func SubmitOrder(c *gin.Context) {
	type request struct {
		CartID    uint   `json:"cart_id" binding:"required"`
		OrderType string `json:"order_type" binding:"required"`
	}
	userID, _ := c.Get("user_id")

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}

	if err := service.SubmitOrder(req.CartID, userID.(uint), req.OrderType); err != nil {
		Respond(c, http.StatusBadRequest, 1, "提交订单失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusBadRequest, 1, "提交订单成功", nil)
}

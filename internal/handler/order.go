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
		PayMethod string `json:"pay_method" binding:"required"`
		Note      string `json:"note"`
	}
	userID, _ := c.Get("user_id")

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}

	order, err := service.SubmitOrder(req.CartID, userID.(uint), req.OrderType, req.Note, req.PayMethod)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "提交订单失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "提交订单成功", gin.H{"order": order})
}

func QueryOrders(c *gin.Context) {
	userID, _ := c.Get("user_id")

	orders, err := service.QueryOrdersByUserID(userID.(uint))
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "获取订单列表失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "获取订单列表成功", gin.H{"orders": orders})
}

func QueryOrderItemsByOrderID(c *gin.Context) {
	type request struct {
		OrderID uint `json:"order_id" binding:"required"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}

	orderItems, err := service.QueryOrderItemsByOrderID(req.OrderID)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "获取订单详情失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "获取订单详情成功", gin.H{"orders": orderItems})
}

func RemoveOrder(c *gin.Context) {
	type request struct {
		OrderID uint `json:"order_id" binding:"required"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}

	err := service.RemoveOrderByOrderID(req.OrderID)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "删除订单失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "删除订单成功", nil)
}

func RefundByOrderID(c *gin.Context) {
	userID, _ := c.Get("user_id")
	type request struct {
		OrderID uint `json:"order_id" binding:"required"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}

	err := service.RefundByOrderID(userID.(uint), req.OrderID)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "退款失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "退款成功", nil)
}

func SubmitRefundRequestByOrderID(c *gin.Context) {
	type request struct {
		OrderID uint `json:"order_id" binding:"required"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}

	err := service.SubmitRefundRequestByOrderID(req.OrderID)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "退款请求失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "退款请求成功", nil)
}

func SubmitCommentByOrderID(c *gin.Context) {
	type request struct {
		OrderID uint   `json:"order_id" binding:"required"`
		Comment string `json:"comment" binding:"required"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}

	err := service.SubmitCommentByOrderID(req.OrderID, req.Comment)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "订单评论失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "订单评论成功", nil)
}

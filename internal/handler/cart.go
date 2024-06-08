package handler

import (
	"net/http"
	"online-ordering-app/internal/service"

	"github.com/gin-gonic/gin"
)

func AddItemToCart(c *gin.Context) {
	type request struct {
		DishID        uint   `json:"dish_id" binding:"required"`
		Specification string `json:"specification"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{
			"field": "dish_id",
			"error": err.Error(),
		})
		return
	}
	userID, _ := c.Get("user_id")

	cartID, err := service.AddToCart(userID.(uint), req.DishID, req.Specification)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "添加详情失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "添加详情成功", gin.H{"cart_id": cartID})
}

func RemoveItemFromCart(c *gin.Context) {
	type request struct {
		CartID        uint   `json:"cart_id" binding:"required"`
		DishID        uint   `json:"dish_id" binding:"required"`
		Specification string `json:"specification"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}

	err := service.RemoveDishFromCartItem(req.CartID, req.DishID, req.Specification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "购物车删除菜品成功", nil)
}

func ClearCart(c *gin.Context) {
	type request struct {
		CartID uint `json:"cart_id" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}

	err := service.ClearCart(req.CartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "购物车已成功清空"})
}

func GetCartInfo(c *gin.Context) {
	type request struct {
		CartID uint `form:"cart_id" binding:"required"`
	}
	var req request
	if err := c.ShouldBindQuery(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{
			"field": "cart_id",
			"error": err.Error(),
		})
		return
	}

	cartItems, err := service.ListCartItemsByCartID(req.CartID)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "获取购物详情失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 2, "获取购物详情成功", gin.H{"cartItems": cartItems})
}

package router

import (
	"online-ordering-app/internal/handler"
	"online-ordering-app/internal/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	r.Static("/assets", "./assets")

	// 小程序接口路由
	apiGroup := r.Group("/api")
	{
		// 身份验证路由组
		authGroup := apiGroup.Group("/auth")
		{
			authGroup.POST("/login", handler.Login)
			authGroup.POST("/register", handler.Register)
		}

		// 用户路由组(小程序)
		userGroup := apiGroup.Group("/user")
		userGroup.Use(middleware.AuthMiddleware())
		{
			userGroup.PUT("/username", handler.UpdateUsername)
			userGroup.PUT("/password", handler.UpdatePassword)
			userGroup.POST("/recharge", handler.RechargeBalance)
			userGroup.POST("/deduct", handler.DeductBalance)
			userGroup.POST("/info", handler.QueryUserInfoByID)
		}

		// 订单路由组（小程序）
		orderGroup := apiGroup.Group("/order")
		orderGroup.Use(middleware.AuthMiddleware())
		{
			orderGroup.POST("/submit", handler.SubmitOrder)
			orderGroup.POST("/list", handler.QueryOrdersByUserID)
			orderGroup.POST("/items", handler.QueryOrderItemsByOrderID)
			orderGroup.POST("/refund_request", handler.SubmitRefundRequestByOrderID)
			orderGroup.POST("/comment", handler.SubmitCommentByOrderID)
			orderGroup.POST("/delete", handler.RemoveOrder)
		}

		// 菜品路由组（小程序）
		dishesGroup := apiGroup.Group("/dish")
		{
			dishesGroup.GET("/list", handler.GetAllDishes)
			dishesGroup.GET("/popular", handler.GetPopularDishes)
			dishesGroup.GET("/info", handler.QueryDishInfoById)
			dishesGroup.GET("/image", handler.QueryDishImageById)
			dishesGroup.GET("/category", handler.GetDishesByCategory)
		}

		// 分类路由组（小程序）
		categoryGroup := apiGroup.Group("/category")
		{
			categoryGroup.GET("/list", handler.GetAllCategories)
		}

		// 购物车路由组(小程序)
		cartGroup := apiGroup.Group("/cart")
		cartGroup.Use(middleware.AuthMiddleware())
		{
			cartGroup.POST("/add", handler.AddItemToCart)
			cartGroup.DELETE("/delete", handler.RemoveItemFromCart)
			cartGroup.DELETE("/clear", handler.ClearCart)

			cartGroup.GET("/info", handler.GetCartInfo)
		}
	}

	// 管理员路由组
	adminGroup := r.Group("/admin")

	adminAuthGroup := adminGroup.Group("/auth")
	{
		adminAuthGroup.POST("/login", handler.Login)
		adminAuthGroup.POST("/register", handler.Register)
	}

	adminGroup.Use(middleware.AuthMiddleware())
	adminGroup.Use(middleware.AdminAuthMiddleware())
	{
		// 管理员-菜品
		adminDishesGroup := adminGroup.Group("/dish")
		{
			adminDishesGroup.POST("/add", handler.AddDish)
			adminDishesGroup.DELETE("/delete", handler.RemoveDish)
			adminDishesGroup.PUT("/update", handler.UpdateDish)
			adminDishesGroup.POST("/list", handler.GetAllDishes)

			adminDishesGroup.GET("/search", handler.QueryDishInfoByKeyword)
			adminDishesGroup.GET("/category", handler.GetDishesByCategory)
		}

		// 管理员-分类
		adminCategoryGroup := adminGroup.Group("/category")
		{
			adminCategoryGroup.POST("/add", handler.AddCategory)
			adminCategoryGroup.DELETE("/delete", handler.RemoveCategory)
			adminCategoryGroup.POST("/list", handler.GetAllCategories)
			adminCategoryGroup.POST("/info", handler.QueryCategoryByID)
		}

		adminOrderGroup := adminGroup.Group("/order")
		{
			adminOrderGroup.POST("/list", handler.QueryAllOrders)
			adminOrderGroup.POST("/finish", handler.FinishOrder)
			adminOrderGroup.POST("/delete", handler.RemoveOrder)
			adminOrderGroup.POST("/refund", handler.RefundByOrderID)
		}
	}

	return r
}

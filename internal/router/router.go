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
		MaxAge: 12 * time.Hour,
	}))

	r.Static("/assets", "./assets")

	// 接口路由
	apiGroup := r.Group("/api")

	// 身份验证路由组
	authGroup := apiGroup.Group("/auth")
	{
		authGroup.POST("/login", handler.Login)
		authGroup.POST("/register", handler.Register)
	}

	// 应用小程序用户路由组
	userGroup := apiGroup.Group("/user")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.PUT("/username", handler.UpdateUsername)
		userGroup.PUT("/password", handler.UpdatePassword)
		userGroup.POST("/recharge", handler.RechargeBalance)
		userGroup.POST("/deduct", handler.DeductBalance)
	}

	// 订单路由组（对外）
	orderGroup := apiGroup.Group("/order")
	orderGroup.Use(middleware.AuthMiddleware())
	{
		orderGroup.POST("/submit", handler.SubmitOrder)
		orderGroup.POST("/list", handler.QueryOrders)
	}

	// 菜品路由组（对外）
	dishesGroup := apiGroup.Group("/dish")
	{
		dishesGroup.GET("/list", handler.GetAllDishes)
		dishesGroup.GET("/search", handler.QueryDishInfoByKeyword)
		dishesGroup.GET("/category", handler.GetDishesByCategory)
		dishesGroup.GET("/popular", handler.GetPopularDishes)
		dishesGroup.GET("/info", handler.QueryDishInfoById)
		dishesGroup.GET("/image", handler.QueryDishImageById)
	}

	// 分类路由组（对外）
	categoryGroup := apiGroup.Group("/category")
	{
		categoryGroup.GET("/list", handler.GetAllCategories)
	}

	// 购物车路由组
	cartGroup := apiGroup.Group("/cart")
	cartGroup.Use(middleware.AuthMiddleware())
	{
		cartGroup.POST("/add", handler.AddItemToCart)
		cartGroup.DELETE("/delete", handler.RemoveItemFromCart)
		cartGroup.DELETE("/clear", handler.ClearCart)

		cartGroup.GET("/info", handler.GetCartInfo)
	}

	// 管理员路由组
	adminGroup := apiGroup.Group("/admin")
	adminGroup.Use(middleware.AuthMiddleware())
	adminGroup.Use(middleware.AdminAuthMiddleware())
	{
		// 管理员对菜品的增删改操作
		adminDishesGroup := adminGroup.Group("/dish")
		{
			adminDishesGroup.POST("/add", handler.AddDish)
			adminDishesGroup.DELETE("/delete", handler.RemoveDish)
			adminDishesGroup.PUT("/update", handler.UpdateDish)
		}

		// 管理员对分类的增删改操作
		adminCategoryGroup := adminGroup.Group("/category")
		{
			adminCategoryGroup.POST("/add", handler.AddCategory)
			adminCategoryGroup.DELETE("/delete", handler.RemoveCategory)
		}
	}

	return r
}

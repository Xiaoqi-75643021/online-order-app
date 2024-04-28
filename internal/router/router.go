package router

import (
	"online-ordering-app/internal/handler"
	"online-ordering-app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

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

	// 菜品路由组（对外）
	dishesGroup := apiGroup.Group("/dish")
	{
		dishesGroup.GET("/list", handler.GetAllDishes)
		dishesGroup.GET("/search", handler.SearchDishes)
		dishesGroup.GET("/category", handler.GetDishesByCategory)
		dishesGroup.GET("/popular", handler.GetPopularDishes)
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
		cartGroup.DELETE("delete", handler.RemoveItemFromCart)

		cartGroup.GET("/list", handler.GetCartInfo)
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

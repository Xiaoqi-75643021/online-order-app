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
	{
		// 身份验证路由组
		authGroup := apiGroup.Group("/auth")
		{
			authGroup.POST("/login", handler.Login)
			authGroup.POST("/register", handler.Register)
		}

		// 用户路由组
		userGroup := apiGroup.Group("/user")
		userGroup.Use(middleware.AuthMiddleware())
		{
			userGroup.PUT("/username", handler.UpdateUsername)
			userGroup.PUT("/password", handler.UpdatePassword)
			userGroup.POST("/recharge", handler.RechargeBalance)
			userGroup.POST("/deduct", handler.DeductBalance)
		}

		// 菜品路由组
		dishesGroup := apiGroup.Group("/dishes")
		{
			dishesGroup.GET("/list", handler.GetAllDishes)
			dishesGroup.GET("/search", handler.SearchDishes)
			dishesGroup.GET("/category", handler.GetDishesByCategory)
			dishesGroup.GET("/popular", handler.GetPopularDishes)
		}
	}

	return r
}

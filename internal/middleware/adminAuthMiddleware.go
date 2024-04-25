package middleware

import (
	"net/http"
	"online-ordering-app/internal/handler"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			handler.Respond(c, http.StatusInternalServerError, 9, "无法获取用户角色", nil)
			c.Abort()
			return
		}

		if role != "admin" {
			handler.Respond(c, http.StatusForbidden, 8, "无管理员权限", nil)
			c.Abort()
			return
		}

		// 如果用户是管理员，继续处理请求
		c.Next()
	}
}

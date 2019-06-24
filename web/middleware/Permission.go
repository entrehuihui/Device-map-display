package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SuperAdministrator 超级管理员权限过滤
func SuperAdministrator() gin.HandlerFunc {
	return func(c *gin.Context) {
		permission := c.GetInt("permisson")
		if permission != 3 {
			userID := c.GetInt("id")
			LogError(fmt.Sprintf("Insufficient permissions SuperAdministrator: userID:%d", userID))
			c.JSON(301, Resqonse{
				Error: 20,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Administrator 管理员权限过滤
func Administrator() gin.HandlerFunc {
	return func(c *gin.Context) {
		permission := c.GetInt("permisson")
		if permission < 2 {
			userID := c.GetInt("id")
			LogError(fmt.Sprintf("Insufficient permissions Administrator: userID:%d", userID))
			c.JSON(301, Resqonse{
				Error: 20,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

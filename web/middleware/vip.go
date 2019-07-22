package middleware

import (
	"../../db"
	"github.com/gin-gonic/gin"
)

// Vip .
func Vip() gin.HandlerFunc {
	return func(c *gin.Context) {
		db.GetRedis()
	}
}

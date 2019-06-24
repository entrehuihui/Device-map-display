package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CheckIP 检测地图服务器IP
func CheckIP() gin.HandlerFunc {
	// ip := viper.GetString("Map.mapurl")
	// ip = strings.Replace(ip, "http://", "", 1)
	// ip = strings.Split(ip, ":")[0]
	return func(c *gin.Context) {
		fmt.Println(c.ClientIP())
		c.Next()
	}
}

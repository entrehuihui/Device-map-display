package middleware

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"../api/service"
	"github.com/gin-gonic/gin"
)

//Verification 登陆检测
func Verification() gin.HandlerFunc {
	return func(c *gin.Context) {
		JWT := c.GetHeader("Authorization")
		if JWT == "" {
			LogError("token contains an invalid number of segments")
			c.JSON(301, 19)
			c.Abort()
			return
		}
		mapClaims, err := service.ParseJWT(JWT)
		if err != nil {
			LogError(err.Error())
			c.JSON(301, 19)
			c.Abort()
			return
		}
		c.Set("id", mapClaims[0])
		c.Set("permisson", mapClaims[1])
		c.Next()
	}
}

// Frequency 频率检测 --- 暂时不用
func Frequency() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// DataLogs ..日志记录
func DataLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		tiBefor := time.Now()
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		c.Next()
		useTime := time.Now().Sub(tiBefor)
		log.Printf("[\t %s]\tMethod[%s]%s [IP]%s [data]%s", useTime, c.Request.Method, c.Request.RequestURI, c.ClientIP(), string(body))
	}
}

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Header("Access-Control-Allow-Methods", "POST, GET,DELETE, PUT, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Accept", "application/json, text/plain, */*")
		c.Header("Origin", "*")
		c.Header("Referer", "*")
		c.Header("User-Agent", "*")
		// //放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			c.Abort()
			return
		}
	}
}

// Corstoo .
func Corstoo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置浏览器不缓存网页请求
		c.Header("cache-control", "no-cache")
		// 处理请求
		c.Next()
	}
}

// http://120.79.252.10:8088/api/vbase/list?status=offline

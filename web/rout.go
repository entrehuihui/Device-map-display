package web

import (
	"log"

	"./api"
	"./middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Run .
// RouterRun .
// @title Swagger Example API
// @version 2.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host petstore.swagger.io
// @BasePath /v2
// RouterRun start server listen
func Run() {
	r := gin.New()
	r.Use(middleware.Cors())

	//图片上传
	r.POST("/upload", api.Upload)

	// swag
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 日志中间件
	r.Use(middleware.DataLogs())

	// 开放图片
	r.Static("/static/", "./static/static/")
	r.Static("/image/", "./static/image/")

	// 设置浏览器不缓存
	// r.Use(middleware.Corstoo())

	// login
	r.PUT("/login", api.Register)
	r.POST("/login", api.Login)
	r.GET("/login", api.CheckUserName)

	//  登陆检测中间件
	r.Use(middleware.Verification())

	// 设备数据
	r.POST("/devicedata", api.SaveDeviceInfos) //数据接收
	r.GET("/devicedata", api.GetDevicesDatas)  //获取数据
	// 设备
	r.GET("/devices", api.GetDevices) // 获取设备列表

	// 管理员权限中间件
	r.Use(middleware.Administrator())
	// 设备
	r.PUT("/devices/status", api.UpdateDevicesStatus) // 更新设备状态
	r.PUT("/devices/user", api.UpdateDevicesUser)     // 更改设备所属权
	r.PUT("/devices/expire", api.UpdateDevicesExpire) // 更改设备过期时间
	r.DELETE("/devices", api.DevicesDel)              // 删除设备
	// 用户
	r.GET("/users", api.GetUsers)                   // 获取用户列表
	r.PUT("/users/expire", api.UpdateDevicesExpire) // 更改用户过期时间
	r.PUT("/users/status", api.UpdateUserStatus)    // 更改用户状态

	// 超级管理员权限
	r.Use(middleware.SuperAdministrator())

	// 启动服务器
	if err := r.Run(viper.GetString("Server.port")); err != nil {
		log.Fatal(err)
	}
}

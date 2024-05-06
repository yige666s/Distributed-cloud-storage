package route

import (
	"Distributed-cloud-storage/handler"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// gin framerwork,获得包含中间件的路由器，包括Logger,Revocery
	router := gin.Default()

	// 处理静态资源
	router.Static("/static/", "./static")

	// 不需要验证的接口
	router.GET("/user/signup", handler.SignUpHandler)
	router.POST("/user/signin", handler.DoSignuphander)

	router.GET("/user/signin", handler.SignInHandler)
	router.POST("/user/signin", handler.DoSignInHandler)

	// 加入中间件，用于校验token的拦截器,此处使用use后所有的handler都会经过拦截器
	router.Use(handler.HTTPInterceptor())
}

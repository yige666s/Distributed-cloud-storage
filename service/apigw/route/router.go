package route

import (
	"Distributed-cloud-storage/service/apigw/handler"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Static("/static/", "./static")
	router.GET("user/signup", handler.SignUpHandler)
	router.GET("user/signup", handler.DoSignuphander)
	return router
}

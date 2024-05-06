package handler

import (
	"Distributed-cloud-storage/common"
	"Distributed-cloud-storage/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// http请求拦截器
func HTTPInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.FormValue("username")
		token := c.Request.FormValue("token")

		if len(username) < 3 || !isToKenValid(token) {
			// token校验失败则直接返回失败提示
			resp := util.NewRespMsg(
				int(common.StatusTokenInvalid),
				"token无效",
				nil,
			)
			c.JSON(http.StatusOK, resp)
			return
		}
		c.Next()
	}
}

// Authorize : http请求拦截器
// func Authorize() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		username := c.Request.FormValue("username")
// 		token := c.Request.FormValue("token")

// 		//验证登录token是否有效
// 		if len(username) < 3 || !isToKenValid(token) {
// 			// w.WriteHeader(http.StatusForbidden)
// 			// token校验失败则跳转到登录页面
// 			c.Abort()
// 			resp := util.NewRespMsg(
// 				int(common.StatusTokenInvalid),
// 				"token无效",
// 				nil,
// 			)
// 			c.JSON(http.StatusOK, resp)
// 			return
// 		}
// 		c.Next()
// 	}
// }

package handler

import (
	"net/http"
)

// http请求拦截器
func HTTPInterceptor(h http.HandleFunc) http.HandleFunc {
	return http.HandleFunc(
		func(w http.ResponseWriter, r *http.Request) {
			r.PostForm()
			username := r.Form.Get("username")
			token := r.Form.Get("token")

			if len(username) < 3 || !isToKenValid(token) {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			h(w, r)
		}
	)
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

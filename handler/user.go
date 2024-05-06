package handler

import (
	"Distributed-cloud-storage/db"
	"Distributed-cloud-storage/util"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	pwd_salt = "*#890" //加盐
)

// 处理Get请求返回注册页面
func SignUpHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signup.html")
}

// 处理POST请求处理注册表单
func DoSignuphander(c *gin.Context) {
	username := c.Request.FormValue(("username"))
	passwd := c.Request.FormValue("password")
	if len(username) < 3 || len(passwd) < 5 {
		c.JSON(http.StatusOK, gin.H{ //  用户明/密码格式不正确
			"msg":  "invalid parameter",
			"code": -1,
		})
		return
	}

	enc_passwd := util.Sha1([]byte(passwd + pwd_salt)) // 对密码进行加盐处理
	suc := db.UserSignup(username, enc_passwd)
	if suc {
		c.JSON(http.StatusOK, gin.H{ //  注册成功
			"msg":  "Signup Succeed",
			"code": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{ //  注册失败
			"msg":  "Signup failed",
			"code": -2,
		})
	}
}

// 返回登陆页面Get请求
func SignInHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signin.html")
}

// 处理POST登录表单请求
func DoSignInHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	enc_passwd := util.Sha1([]byte(password + pwd_salt))
	// 1. 校验用户名和密码
	pwdChecked := db.UserSignIn(username, enc_passwd)
	if !pwdChecked {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "login failed",
			"code": -1,
		})
		return
	}
	// 2. 生成访问凭证(token)
	token := GenToken(username)
	upRes := db.UpdateToken(username, token)
	if !upRes {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "login failed",
			"code": -2,
		})
		return
	}
	// 3. 登陆成功后重定向至首页,具体重定向操作由客户端实现
	// w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}
	c.Data(http.StatusOK, "application/json", resp.JSONBytes())
}

func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// 1.解析请求参数
	r.ParseForm()
	username := r.Form.Get("username")
	// token := r.Form.Get("token")
	// // 2. 验证Token是否有效
	// isValidToken := isToKenValid(token)
	// if !isValidToken {
	// 	w.WriteHeader(http.StatusForbidden) // 403错误
	// 	return
	// 3. 查询用户信息
	user, err := db.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// 4. 组装并响应用户数据
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	w.Write(resp.JSONBytes())
}

// 验证token是否有效
func isToKenValid(token string) bool {
	if len(token) != 40 {
		return false
	}
	// 判断Token时效性
	// 从数据库表tbl_user_token查询username对应token
	// 对比两个token是否一致
	return true
}

// 生成40位Token
func GenToken(username string) string {
	// md5(username +timestamp +token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

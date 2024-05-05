package handler

import (
	"Distributed-cloud-storage/db"
	"Distributed-cloud-storage/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	pwd_salt = "*#890" //加盐
)

// 处理用户注册请求
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet { // Get请求，返回注册页面
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm() //POST请求，用户注册
	username := r.Form.Get(("username"))
	passwd := r.Form.Get("password")
	if len(username) < 3 || len(passwd) < 5 {
		w.Write([]byte("invalid paprameter")) // 用户明/密码格式不正确
		return
	}

	enc_passwd := util.Sha1([]byte(passwd + pwd_salt)) // 对密码进行加盐处理
	suc := db.UserSignup(username, enc_passwd)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}
}

// 登录接口
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	enc_passwd := util.Sha1([]byte(password + pwd_salt))
	// 1. 校验用户名和密码
	pwdChecked := db.UserSignIn(username, enc_passwd)
	if !pwdChecked {
		w.Write([]byte("FAILED"))
		return
	}
	// 2. 生成访问凭证(token)
	token := GenToken(username)
	upRes := db.UpdateToken(username, token)
	if !upRes {
		w.Write([]byte("FAILED"))
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
			Location: "http://" + r.Host + "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}
	w.Write(resp.JSONBytes()) // 以Json形式返回数据
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

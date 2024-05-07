package handler

import (
	"Distributed-cloud-storage/service/account/proto/user"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"micro.dev/v4/cmd/protoc-gen-micro/plugin/micro"
)

var (
	userCli user.UserService
)

func init() {
	service := micro.NewService()
	//初始化解析命令行参数
	service.Init()

	// 初始化一个rpcclient
	userCli = user.NewUserService("go.micro.servic.user", service.Client())
}

func SignUpHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signup.html")
}

// 将用户http请求转换为RPC
func DoSignuphander(c *gin.Context) {
	username := c.Request.FormValue(("username"))
	passwd := c.Request.FormValue("password")

	resp, err := userCli.SignUp(context.TODO(), &user.ReqSignup{Username: username, Password: passwd})
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Message,
	})
}

package main

import (
	"Distributed-cloud-storage/service/account/proto/user"
	"time"

	micro "github.com/micro/go-micro"
)

func main() {
	// 创建一个服务
	service := micro.NewService(micro.Name("go.micro.service.user"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5))
	service.Init()
	user.RegisterUserServiceHandler(service.Service())
	// TODO:没写完
}

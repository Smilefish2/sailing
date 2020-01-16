package main

import (
	"github.com/Smilefish0/sailing/configer"
	_ "github.com/Smilefish0/sailing/starter"
	"github.com/Smilefish0/sailing/user/handler"
	srv "github.com/Smilefish0/sailing/user/service"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	pb "github.com/Smilefish0/sailing/user/proto"
)

func main() {

	// New Service   新建服务
	service := micro.NewService(
		micro.Name(configer.AppConfig().GetName()),
		micro.Version("latest"),
	)

	// Initialise service  初始化服务
	service.Init(
		micro.Action(func(c *cli.Context) {
			srv.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// Register Handler   注册服务
	// service.RegisterUserHandler(service.Server(), new(handler.Service))
	pb.RegisterUserHandler(service.Server(), new(handler.Service))

	// Run service    启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

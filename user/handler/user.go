package handler

import (
	"context"
	userProto "github.com/Smilefish0/sailing/user/proto/user"
	"github.com/Smilefish0/sailing/user/service"
	"github.com/micro/go-micro/util/log"
)

type Service struct{}

var (
	userService service.UserService
)

// Init 初始化handler
func Init() {
	var err error
	userService, err = service.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

// QueryUserByName 通过参数中的名字返回用户
func (e *Service) QueryUserByName(ctx context.Context, req *userProto.Request, rsp *userProto.Response) error {
	user, err := userService.QueryUserByName(req.Username)
	if err != nil {
		rsp.Success = false
		rsp.Error = &userProto.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return nil
	}

	rsp.User = user
	rsp.Success = true
	return nil
}

package service

import (
	"fmt"
	"github.com/Smilefish2/sailing/starter"
	userProto "github.com/Smilefish2/sailing/user/proto/user"
	"sync"
)

var (
	s *userService
	m sync.RWMutex
)

// service 服务
type userService struct {
}

// Service 用户服务类
type UserService interface {
	// QueryUserByName 根据用户名获取用户
	QueryUserByName(userName string) (ret *userProto.User, err error)
}

// GetService 获取服务类
func GetService() (*userService, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] UserService 未初始化")
	}
	return s, nil
}

// Init 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	s = &userService{}
}

func (s *userService) QueryUserByName(userName string) (ret *userProto.User, err error) {
	queryString := `SELECT uid, username, password FROM user WHERE username = ?`

	// 获取数据库
	db := starter.DB()

	ret = &userProto.User{}

	// 查询
	db.Raw(queryString, userName).Scan(&ret)
	//errs := db.GetErrors()
	//spew.Dump(ret)
	//if errs != nil {
	//	log.Logf("[QueryUserByName] 查询数据失败，err：%s", errs[0])
	//	return
	//}

	//ret = &proto.User{
	//	m.ID,
	//	m.Username,
	//	m.Password,
	//}
	return
}

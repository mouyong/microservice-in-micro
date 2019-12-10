package handler

import (
	"context"
	"github.com/micro/go-micro/util/log"
	s "microservice-in-micro/user-srv/proto/user"
	us "microservice-in-micro/user-srv/model/user"
)

var (
	userService us.Service
)

func Init()  {
	var err error

	userService, err = us.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

type User struct{}


func (u *User) QueryUserByName(ctx context.Context, req *s.Request, resp *s.Response) error {
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		resp.Success = false
		resp.Error = &s.Error{
			Code:                 500,
			Detail:               err.Error(),
		}
		return err
	}
	resp.User = user
	resp.Success = true
	return nil
}



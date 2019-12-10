package handler

import (
	"context"
	"github.com/micro/go-micro/util/log"
	"microservice-in-micro/auth/model/access"
	"microservice-in-micro/auth/proto/auth"
)

var (
	accessService access.Service
)

func Init()  {
	var err error
	accessService, err = access.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化 Handler 错误，%s", err)
		return
	}
}

type Service struct{}

func (s *Service) MakeAccessToken(ctx context.Context, req *mu_micro_book_srv_auth.Request, rsp *mu_micro_book_srv_auth.Response) error {
	log.Log("[MakeAccessToken] 收到创建token请求")

	token, err := accessService.MakeAccessToken(&access.Subject{
		ID: string(req.UserId),
		Name: req.UserName,
	})
	if err != nil {
		rsp.Error = &mu_micro_book_srv_auth.Error{
			Detail: err.Error(),
		}

		log.Logf("[MakeAccessToken] token生成失败，err：%s", err)
		return err
	}

	rsp.Token = token
	return nil
}

func (s *Service) DelUserAccessToken(ctx context.Context, req *mu_micro_book_srv_auth.Request, rsp *mu_micro_book_srv_auth.Response) error {
	log.Log("[DelUserAccessToken] 清除用户token")
	err := accessService.DelUserAccessToken(req.Token)
	if err != nil {
		rsp.Error = &mu_micro_book_srv_auth.Error{
			Detail: err.Error(),
		}

		log.Logf("[DelUserAccessToken] 清除用户token失败，err：%s", err)
		return err
	}

	return nil
}
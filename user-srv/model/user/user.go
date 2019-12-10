package user

import (
	"fmt"
	"sync"
	"microservice-in-micro/user-srv/proto/user"
)

var (
	s *service
	m sync.RWMutex
)

type service struct {}

type Service interface {
	QueryUserByName(userName string) (ret *mu_micro_book_srv_user.User, err error)
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化 ")
	}

	return s, nil
}

func Init()  {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	s = &service{}
}
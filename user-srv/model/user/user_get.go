package user

import (
	"github.com/micro/go-micro/util/log"
	"microservice-in-micro/basic/db"
	"microservice-in-micro/user-srv/proto/user"
)

func (u service) QueryUserByName(userName string) (ret *mu_micro_book_srv_user.User, err error) {
	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`

	o := db.GetDb()

	ret = &mu_micro_book_srv_user.User{}

	err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}
	return
}
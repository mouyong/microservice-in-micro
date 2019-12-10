package db

import (
	"database/sql"
	"fmt"
	"github.com/micro/go-micro/util/log"
	"sync"
	"microservice-in-micro/basic/config"
)

var (
	inited bool
	m sync.RWMutex
	mysqlDB *sql.DB
)

func Init()  {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] db 已经初始化过")
		log.Logf(err.Error())
		return
	}

	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}
	
	inited = true
}

func GetDb() *sql.DB {
	if mysqlDB == nil {
		log.Fatal("[GetDb] mysqlDB 尚未初始化")
	}
	return mysqlDB
}
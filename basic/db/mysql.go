package db

import (
	"database/sql"
	"github.com/micro/go-micro/util/log"
	"microservice-in-micro/basic/config"
)

func initMysql()  {
	var err error

	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetUrl())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Logf("[initMysql] 初始化数据库成功")
}

package basic

import (
	"microservice-in-micro/basic/config"
	"microservice-in-micro/basic/db"
	"microservice-in-micro/basic/redis"
)

func Init()  {
	config.Init()
	db.Init()
	redis.Init()
}
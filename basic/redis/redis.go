package redis

import (
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/util/log"
	"microservice-in-micro/basic/config"
	"sync"
)

var (
	client *redis.Client
	m sync.RWMutex
	inited bool
)

func Init()  {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Log("已经初始化过Redis...")
		return
	}

	redisConfig := config.GetRedisConfig()

	if redisConfig != nil && redisConfig.GetEnabled() {
		log.Log("初始化Redis...")

		if redisConfig.GetSentinelConfig() != nil && redisConfig.GetSentinelConfig().GetEnabled() {
			log.Log("初始化Redis，哨兵模式...")
			initSentinel(redisConfig)
		} else {
			log.Log("初始化Redis，普通模式...")
			initSingle(redisConfig)
		}

		log.Log("初始化Redis，检测连接...")

		pong, err := client.Ping().Result()
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Log("初始化Redis，检测连接Ping.")
		log.Log("初始化Redis，检测连接Ping..")
		log.Logf("初始化Redis，检测连接Ping... %s", pong)
	}

	inited = true
}

func GetRedis() *redis.Client {
	if client == nil {
		log.Fatal("[GetRedis] redis 尚未初始化")
	}
	return client
}

func initSentinel(redisConfig config.RedisConfig)  {
	client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    redisConfig.GetSentinelConfig().GetMaster(),
		SentinelAddrs: redisConfig.GetSentinelConfig().GetNodes(),
		DB:            redisConfig.GetDBNum(),
		Password:      redisConfig.GetPassword(),
	})
}

func initSingle(redisConfig config.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.GetConn(),
		Password: redisConfig.GetPassword(), // no password set
		DB:       redisConfig.GetDBNum(),    // use default DB
	})
}
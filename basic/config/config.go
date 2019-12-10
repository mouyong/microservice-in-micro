package config

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/config"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	err error
)

var (
	defaultRootPath = "app"
	defaultConfigFilePrefix = "application-"
	profiles defaultProfiles
	etcdConfig defaultEtcdConfig
	mysqlConfig defaultMysqlConfig
	jwtConfig defaultJwtConfig
	redisConfig defaultRedisConfig
	m sync.RWMutex
	inited bool
)

func Init()  {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Logf("[Init] 配置已经初始化过")
		return
	}

	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("./"+string(filepath.Separator))))
	pt := filepath.Join(appPath, "conf")
	os.Chdir(pt)

	if err = config.Load(file.NewSource(file.WithPath(pt+"/application.yml"))); err != nil {
		panic(err)
	}

	if err = config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}

	log.Logf("[Init] 加载配置文件：path: %s, %+v\n", pt+"/application.yml", profiles)
	if len(profiles.GetInclude()) > 0 {
		include := strings.Split(profiles.GetInclude(), ",")

		sources := make([]source.Source, len(include))
		for i:=0; i<len(include); i++ {
			filePath := pt+ string(filepath.Separator) +defaultConfigFilePrefix+strings.TrimSpace(include[i])+".yml"

			log.Logf("[Init] 加载配置文件：path: %s\n", filePath)

			sources[i] = file.NewSource(file.WithPath(filePath))
		}

		if err = config.Load(sources...); err != nil {
			panic(err)
		}
	}

	config.Get(defaultRootPath, "etcd").Scan(&etcdConfig)
	config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)
	config.Get(defaultRootPath, "jwt").Scan(&jwtConfig)
	config.Get(defaultRootPath, "redis").Scan(&redisConfig)

	inited = true
}

func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

func GetEtcdConfig() (ret EtcdConfig) {
	return etcdConfig
}

func GetJwtConfig() (ret JwtConfig) {
	return jwtConfig
}

func GetRedisConfig() (ret RedisConfig) {
	return redisConfig
}
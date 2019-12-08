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
	m sync.RWMutex
	inited bool
	sp = string(filepath.Separator)
)

func Init()  {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Logf("[Init] 配置已经初始化过")
		return
	}

	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("."+sp, sp)))
	pt := filepath.Join(appPath, "conf")
	os.Chdir(pt)

	if err = config.Load(file.NewSource(file.WithPath(pt+sp+"application.yml"))); err != nil {
		panic(err)
	}

	if err = config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}

	log.Logf("[Init] 加载配置文件：path: %s, %+v\n", pt+sp+"application.yml", profiles)
	if len(profiles.GetInclude()) > 0 {
		include := strings.Split(profiles.GetInclude(), ",")

		sources := make([]source.Source, len(include))
		for i:=0; i<len(include); i++ {
			filePath := pt+sp+defaultConfigFilePrefix+strings.TrimSpace(include[i])+".yml"

			log.Logf("[Init] 加载配置文件：path: %s\n", filePath)

			sources[i] = file.NewSource(file.WithPath(filePath))
		}

		if err = config.Load(sources...); err != nil {
			panic(err)
		}
	}

	config.Get(defaultRootPath, "etcd").Scan(&etcdConfig)
	config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)

	inited = true
}

func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

func GetEtcdConfig() (ret EtcdConfig) {
	return etcdConfig
}
package config

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	err                     error
	defaultRootPath         = "app"
	defaultConfigFilePrefix = "application-"
	consulConfig            defaultConsulConfig
	mysqlConfig             defaultMysqlConfig
	profiles                defaultProfiles
	m                       sync.RWMutex
	inited                  bool
	sp                      = string(filepath.Separator)
)

func Init() {
	m.Lock()
	defer m.Unlock()
	if inited {
		log.Logf("[Init]配置已经初始化过")
		return
	}
	//加载conf下的yml文件
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("."+sp, sp)))
	pt := filepath.Join(appPath, "conf")
	_ = os.Chdir(appPath)
	// 找到application.yml文件并加载，生成source.Source实例并加入到DefaultConfig中
	if err = config.Load(file.NewSource(file.WithPath(pt + sp + "application.yml"))); err != nil {
		panic(err)
	}
	// 从上面加载的配置中获取一个配置值app-profiles-下的键值对，然后存入profiles
	if err = config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}
	log.Logf("[Init]加载配置文件：path: %s, %+v\n", pt+sp+"application.yml", profiles)
	//开始导入新文件，就是把db和consul两个后缀的yml文件的信息读到source.Source实例再加入到DefaultConfig中
	if len(profiles.GetInclude()) > 0 {
		include := strings.Split(profiles.GetInclude(), ",")
		sources := make([]source.Source, len(include))
		for i := 0; i < len(include); i++ {
			filePath := pt + string(filepath.Separator) + defaultConfigFilePrefix + strings.TrimSpace(include[i]) + ".yml"
			log.Logf("[Init]加载配置文件：path: %s\n", filePath)
			sources[i] = file.NewSource(file.WithPath(filePath))
		}
		//记载include的文件
		if err = config.Load(sources...); err != nil {
			panic(err)
		}
	}
	// 赋值
	_ = config.Get(defaultRootPath, "consul").Scan(&consulConfig)
	_ = config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)

	// 标记已经初始化
	inited = true
}

// GetMysqlConfig 获取mysql配置
func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

// GetConsulConfig 获取Consul配置
func GetConsulConfig() (ret ConsulConfig) {
	return consulConfig
}

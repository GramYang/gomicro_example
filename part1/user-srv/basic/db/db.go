package db

import (
	"database/sql"
	"fmt"
	"github.com/micro/go-micro/util/log"
	"gomicro_example/part1/user-srv/basic/config"
	"sync"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

func Init() {
	m.Lock()
	defer m.Unlock()
	var err error
	if inited {
		err = fmt.Errorf("[Init] db 已经初始化过了")
		log.Logf(err.Error())
		return
	}
	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}
	inited = true
}

func GetDB() *sql.DB {
	return mysqlDB
}

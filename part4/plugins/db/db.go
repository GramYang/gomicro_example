package db

import (
	"database/sql"
	"fmt"
	"github.com/micro/go-micro/util/log"
	"gomicro_example/part4/basic"
	"sync"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

func init() {
	basic.Register(initDB)
}

// initDB 初始化数据库
func initDB() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[initDB] db 已经初始化过")
		log.Logf(err.Error())
		return
	}

	initMysql()

	inited = true
}

// GetDB 获取db
func GetDB() *sql.DB {
	return mysqlDB
}

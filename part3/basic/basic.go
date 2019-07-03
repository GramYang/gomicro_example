package basic

import (
	"gomicro_example/part3/basic/config"
	"gomicro_example/part3/basic/db"
	"gomicro_example/part3/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}

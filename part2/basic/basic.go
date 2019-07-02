package basic

import (
	"gomicro_example/part2/basic/config"
	"gomicro_example/part2/basic/db"
	"gomicro_example/part2/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}

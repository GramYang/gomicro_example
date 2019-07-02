package basic

import (
	"gomicro_example/part1/user-srv/basic/config"
	"gomicro_example/part1/user-srv/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}

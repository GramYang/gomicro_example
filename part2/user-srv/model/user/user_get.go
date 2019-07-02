package user

import (
	"github.com/micro/go-micro/util/log"
	proto "gomicro_example/part1/proto/user"
	"gomicro_example/part1/user-srv/basic/db"
)

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
	queryString := `select user_id, user_name, pwd from user where user_name = ?`
	o := db.GetDB()
	ret = &proto.User{}
	err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}
	return
}

package user

import (
	"github.com/micro/go-micro/util/log"
	"gomicro_example/part6/plugins/db"
	proto "gomicro_example/part6/user-srv/proto/user"
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

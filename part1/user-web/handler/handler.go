package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	us "gomicro_example/part1/proto/user"
	"net/http"
	"time"
)

var (
	serviceClient us.UserService
)

type Error struct {
	Code   string
	Detail string
}

func Init() {
	serviceClient = us.NewUserService("mu.micro.book.srv.user", client.DefaultClient)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Logf("只能是POST请求")
		http.Error(w, "非法请求", 400)
		return
	}
	_ = r.ParseForm()
	rsp, err := serviceClient.QueryUserByName(context.TODO(), &us.Request{
		UserName: r.Form.Get("userName"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}
	if rsp.User.Pwd == r.Form.Get("pwd") {
		response["success"] = rsp.Success
		rsp.User.Pwd = ""
		response["data"] = rsp.User
	} else {
		response["success"] = false
		response["error"] = &Error{
			Detail: "密码错误",
		}
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

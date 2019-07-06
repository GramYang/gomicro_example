package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/util/log"
	"net/http"
	"strconv"
	"time"

	"github.com/micro/go-micro/client"
	auth "gomicro_example/part4/auth/proto/auth"
	invS "gomicro_example/part4/inventory-srv/proto/inventory"
	order "gomicro_example/part4/order-srv/proto/order"
	"gomicro_example/part4/plugins/session"
)

var (
	serviceClient order.OrdersService
	authClient    auth.Service
	invClient     invS.InventoryService
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = order.NewOrdersService("mu.micro.book.srv.orders", client.DefaultClient)
	authClient = auth.NewService("mu.micro.book.srv.auth", client.DefaultClient)
}

// New 新增订单入口
func New(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	_ = r.ParseForm()
	bookId, _ := strconv.ParseInt(r.Form.Get("bookId"), 10, 10)

	// 返回结果
	response := map[string]interface{}{}

	// 调用后台服务
	rsp, err := serviceClient.New(context.TODO(), &order.Request{
		BookId: bookId,
		UserId: session.GetSession(w, r).Values["userId"].(int64),
	})

	// 返回结果
	response["ref"] = time.Now().UnixNano()
	if err != nil {
		response["success"] = false
		response["error"] = Error{
			Detail: err.Error(),
		}
	} else {
		response["success"] = true
		response["orderId"] = rsp.Order.Id
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

package main

import (
	"fmt"
	"im-chat-server/controllers"
	"net/http"
)

func main() {
	go http.ListenAndServe(":5001", controllers.NewHomeController())
	fmt.Println("服务器已启动,在5001端口监听")
	select {}
}

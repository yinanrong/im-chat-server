package main

import (
	"im-chat-server/controllers"
	"net/http"
)

func main() {
	http.ListenAndServe(":5001", controllers.NewHomeController())
}

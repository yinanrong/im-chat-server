package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"im-chat-server/models"
	"io/ioutil"
	"net/http"
)

type HomeController struct {
	controller
}

func NewHomeController() *HomeController {
	h := new(HomeController)
	if h.mux == nil {
		h.mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	}
	h.mux["/"] = h.callback
	return h
}

func (c *HomeController) callback(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	defer r.Body.Close()
	if byteData, err := ioutil.ReadAll(r.Body); err == nil {
		json.Unmarshal(byteData, &message)
		logs.Info("收到消息:", message)
	} else {
		logs.Error(err.Error())
	}
	c.OK(w, message)
}

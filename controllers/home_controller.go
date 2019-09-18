package controllers

import (
	"im-chat-server/models"
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

	c.OK(w, models.Message{})
}

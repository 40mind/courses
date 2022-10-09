package controller

import (
	"courses/service"
	"github.com/jimlawless/whereami"
	"html/template"
	"log"
	"net/http"
)

type Controller struct {
	Service		*service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		Service: service,
	}
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Printf("controller error: %s: %s", err.Error(), whereami.WhereAmI())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("controller error: %s: %s", err.Error(), whereami.WhereAmI())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

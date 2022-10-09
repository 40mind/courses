package router

import (
	"courses/presentation/controller"
	"github.com/gorilla/mux"
)

func NewRouter(c *controller.Controller) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", c.Index).Methods("GET")
	return r
}

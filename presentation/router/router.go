package router

import (
	"courses/presentation/controller"
	"github.com/gorilla/mux"
)

func NewRouter(c *controller.Controller) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/tech/info", c.TechInfo).Methods("GET")
	r.HandleFunc("/admin", c.AdminHome).Methods("GET")
	r.HandleFunc("/course/{id}", c.CoursePage).Methods("GET")
	r.HandleFunc("/course/{id}", c.CreateStudent).Methods("POST")
	r.HandleFunc("/", c.HomePage).Methods("GET")
	return r
}

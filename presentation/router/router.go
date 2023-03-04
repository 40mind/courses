package router

import (
	"courses/presentation/controller"
	"github.com/gorilla/mux"
)

func NewRouter(c *controller.Controller) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/admin/courses/{id}", c.AdminUpdateCourse).Methods("PATCH")
	r.HandleFunc("/admin/courses/{id}", c.AdminDeleteCourse).Methods("DELETE")
	r.HandleFunc("/admin", c.AdminHome).Methods("GET")
	r.HandleFunc("/{id}", c.CoursePage).Methods("GET")
	r.HandleFunc("/{id}", c.CreateStudent).Methods("POST")
	r.HandleFunc("/", c.HomePage).Methods("GET")
	return r
}

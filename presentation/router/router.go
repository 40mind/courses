package router

import (
	"courses/presentation/controller"
	"github.com/gorilla/mux"
)

func NewRouter(c *controller.Controller) *mux.Router {
	r := mux.NewRouter()
	addAdminRoutes(r, c)
	addUserRoutes(r, c)
	addTechRoutes(r, c)
	return r
}

func addAdminRoutes(r *mux.Router, c *controller.Controller) {
	r.HandleFunc("/admin", c.AdminHome).Methods("GET")

	r.HandleFunc("/admin/directions", c.AdminDirections).Methods("GET")
	r.HandleFunc("/admin/directions", c.AdminCreateDirection).Methods("POST")
	r.HandleFunc("/admin/directions/{id}", c.AdminDirection).Methods("GET")
	r.HandleFunc("/admin/directions/{id}", c.AdminUpdateDirection).Methods("PATCH")
	r.HandleFunc("/admin/directions/{id}", c.AdminDeleteDirection).Methods("DELETE")

	r.HandleFunc("/admin/courses", c.AdminCourses).Methods("GET")
	r.HandleFunc("/admin/courses", c.AdminCreateCourse).Methods("POST")
	r.HandleFunc("/admin/courses/{id}", c.AdminCourse).Methods("GET")
	r.HandleFunc("/admin/courses/{id}", c.AdminUpdateCourse).Methods("PATCH")
	r.HandleFunc("/admin/courses/{id}", c.AdminDeleteCourse).Methods("DELETE")

	r.HandleFunc("/admin/students", c.AdminStudents).Methods("GET")
	r.HandleFunc("/admin/students", c.AdminCreateStudent).Methods("POST")
	r.HandleFunc("/admin/students/{id}", c.AdminStudent).Methods("GET")
	r.HandleFunc("/admin/students/{id}", c.AdminUpdateStudent).Methods("PATCH")
	r.HandleFunc("/admin/students/{id}", c.AdminDeleteStudent).Methods("DELETE")
}

func addUserRoutes(r *mux.Router, c *controller.Controller) {
	r.HandleFunc("/course/{id}", c.CoursePage).Methods("GET")
	r.HandleFunc("/course/{id}", c.CreateStudent).Methods("POST")
	r.HandleFunc("/", c.HomePage).Methods("GET")
}

func addTechRoutes(r *mux.Router, c *controller.Controller) {
	r.HandleFunc("/tech/info", c.TechInfo).Methods("GET")
}

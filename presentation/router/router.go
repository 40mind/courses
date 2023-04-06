package router

import (
	"courses/middlewares"
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
	r.HandleFunc("/admin", middlewares.AuthMiddleware(c.AdminHome, c.Store)).Methods("GET")

	r.HandleFunc("/admin/login", c.AdminLogIn).Methods("POST")
	r.HandleFunc("/admin/logout", c.AdminLogOut).Methods("POST")
	r.HandleFunc("/admin/admins", middlewares.AuthMiddleware(c.AdminGetAdmins, c.Store)).Methods("GET")
	r.HandleFunc("/admin/admins", middlewares.AuthMiddleware(c.AdminCreateAdmin, c.Store)).Methods("POST")
	r.HandleFunc("/admin/admins/{id}", middlewares.AuthMiddleware(c.AdminGetAdmin, c.Store)).Methods("GET")
	r.HandleFunc("/admin/admins/{id}", middlewares.AuthMiddleware(c.DeleteAdmin, c.Store)).Methods("DELETE")

	r.HandleFunc("/admin/directions", middlewares.AuthMiddleware(c.AdminDirections, c.Store)).Methods("GET")
	r.HandleFunc("/admin/directions", middlewares.AuthMiddleware(c.AdminCreateDirection, c.Store)).Methods("POST")
	r.HandleFunc("/admin/directions/{id}", middlewares.AuthMiddleware(c.AdminDirection, c.Store)).Methods("GET")
	r.HandleFunc("/admin/directions/{id}", middlewares.AuthMiddleware(c.AdminUpdateDirection, c.Store)).Methods("PATCH")
	r.HandleFunc("/admin/directions/{id}", middlewares.AuthMiddleware(c.AdminDeleteDirection, c.Store)).Methods("DELETE")

	r.HandleFunc("/admin/courses", middlewares.AuthMiddleware(c.AdminCourses, c.Store)).Methods("GET")
	r.HandleFunc("/admin/courses", middlewares.AuthMiddleware(c.AdminCreateCourse, c.Store)).Methods("POST")
	r.HandleFunc("/admin/courses/{id}", middlewares.AuthMiddleware(c.AdminCourse, c.Store)).Methods("GET")
	r.HandleFunc("/admin/courses/{id}", middlewares.AuthMiddleware(c.AdminUpdateCourse, c.Store)).Methods("PATCH")
	r.HandleFunc("/admin/courses/{id}", middlewares.AuthMiddleware(c.AdminDeleteCourse, c.Store)).Methods("DELETE")

	r.HandleFunc("/admin/students", middlewares.AuthMiddleware(c.AdminStudents, c.Store)).Methods("GET")
	r.HandleFunc("/admin/students", middlewares.AuthMiddleware(c.AdminCreateStudent, c.Store)).Methods("POST")
	r.HandleFunc("/admin/students/{id}", middlewares.AuthMiddleware(c.AdminStudent, c.Store)).Methods("GET")
	r.HandleFunc("/admin/students/{id}", middlewares.AuthMiddleware(c.AdminUpdateStudent, c.Store)).Methods("PATCH")
	r.HandleFunc("/admin/students/{id}", middlewares.AuthMiddleware(c.AdminDeleteStudent, c.Store)).Methods("DELETE")
}

func addUserRoutes(r *mux.Router, c *controller.Controller) {
	r.HandleFunc("/course/{id}", c.CoursePage).Methods("GET")
	r.HandleFunc("/course/{id}", c.CreateStudent).Methods("POST")
	r.HandleFunc("/payment/create/{id}", c.CreatePayment).Methods("GET")
	r.HandleFunc("/payment/confirm/{id}", c.ConfirmPayment).Methods("GET")
	r.HandleFunc("/", c.HomePage).Methods("GET")
}

func addTechRoutes(r *mux.Router, c *controller.Controller) {
	r.HandleFunc("/tech/info", c.TechInfo).Methods("GET")
}

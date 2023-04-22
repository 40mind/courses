package router

import (
	"courses/middlewares"
	"courses/presentation/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(c *controller.Controller) *mux.Router {
	r := mux.NewRouter()
	addAdminRoutes(r, c)
	addUserRoutes(r, c)
	addTechRoutes(r, c)
	addFrontRoutes(r)
	return r
}

func addAdminRoutes(r *mux.Router, c *controller.Controller) {
	r.HandleFunc("/api/v1/admin", middlewares.AuthMiddleware(c.AdminHome, c.Store)).Methods("GET")

	r.HandleFunc("/api/v1/admin/login", c.AdminLogIn).Methods("POST")
	r.HandleFunc("/api/v1/admin/logout", c.AdminLogOut).Methods("POST")
	r.HandleFunc("/api/v1/admin/admins", middlewares.AuthMiddleware(c.AdminGetAdmins, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/admins", middlewares.AuthMiddleware(c.AdminCreateAdmin, c.Store)).Methods("POST")
	r.HandleFunc("/api/v1/admin/admins/{id}", middlewares.AuthMiddleware(c.AdminGetAdmin, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/admins/{id}", middlewares.AuthMiddleware(c.DeleteAdmin, c.Store)).Methods("DELETE")

	r.HandleFunc("/api/v1/admin/directions", middlewares.AuthMiddleware(c.AdminDirections, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/directions", middlewares.AuthMiddleware(c.AdminCreateDirection, c.Store)).Methods("POST")
	r.HandleFunc("/api/v1/admin/directions/{id}", middlewares.AuthMiddleware(c.AdminDirection, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/directions/{id}", middlewares.AuthMiddleware(c.AdminUpdateDirection, c.Store)).Methods("PATCH")
	r.HandleFunc("/api/v1/admin/directions/{id}", middlewares.AuthMiddleware(c.AdminDeleteDirection, c.Store)).Methods("DELETE")

	r.HandleFunc("/api/v1/admin/courses", middlewares.AuthMiddleware(c.AdminCourses, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/courses", middlewares.AuthMiddleware(c.AdminCreateCourse, c.Store)).Methods("POST")
	r.HandleFunc("/api/v1/admin/courses/{id}", middlewares.AuthMiddleware(c.AdminCourse, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/courses/{id}", middlewares.AuthMiddleware(c.AdminUpdateCourse, c.Store)).Methods("PATCH")
	r.HandleFunc("/api/v1/admin/courses/{id}", middlewares.AuthMiddleware(c.AdminDeleteCourse, c.Store)).Methods("DELETE")

	r.HandleFunc("/api/v1/admin/students", middlewares.AuthMiddleware(c.AdminStudents, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/students", middlewares.AuthMiddleware(c.AdminCreateStudent, c.Store)).Methods("POST")
	r.HandleFunc("/api/v1/admin/students/{id}", middlewares.AuthMiddleware(c.AdminStudent, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/students/{id}", middlewares.AuthMiddleware(c.AdminUpdateStudent, c.Store)).Methods("PATCH")
	r.HandleFunc("/api/v1/admin/students/{id}", middlewares.AuthMiddleware(c.AdminDeleteStudent, c.Store)).Methods("DELETE")
}

func addUserRoutes(r *mux.Router, c *controller.Controller) {
	r.HandleFunc("/api/v1/course/{id}", c.CoursePage).Methods("GET")
	r.HandleFunc("/api/v1/course/{id}", c.CreateStudent).Methods("POST")
	r.HandleFunc("/api/v1/payment/create/{id}", c.CreatePayment).Methods("POST")
	r.HandleFunc("/api/v1/payment/confirm/{id}", c.ConfirmPayment).Methods("POST")
	r.HandleFunc("/api/v1", c.HomePage).Methods("GET")
}

func addTechRoutes(r *mux.Router, c *controller.Controller) {
	r.HandleFunc("/api/v1/tech/info", c.TechInfo).Methods("GET")
}

func addFrontRoutes(r *mux.Router) {
	fs := http.FileServer(http.Dir("./front/"))
	r.PathPrefix("").Handler(fs)
}

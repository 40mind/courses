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
	addEditorRoutes(r, c)
	addUserRoutes(r, c)
	addTechRoutes(r, c)
	addFrontRoutes(r)
	return r
}

func addAdminRoutes(r *mux.Router, c *controller.Controller) {
	r.HandleFunc("/api/v1/admin", middlewares.AuthAdminMiddleware(c.AdminHome, c.Store)).Methods("GET")

	r.HandleFunc("/api/v1/admin/login", c.AdminLogIn).Methods("POST")
	r.HandleFunc("/api/v1/admin/logout", c.AdminLogOut).Methods("POST")

	r.HandleFunc("/api/v1/admin/admins", middlewares.AuthAdminMiddleware(c.GetAdmins, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/admins", middlewares.AuthAdminMiddleware(c.CreateAdmin, c.Store)).Methods("POST")
	r.HandleFunc("/api/v1/admin/admins/{id}", middlewares.AuthAdminMiddleware(c.GetAdmin, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/admins/{id}", middlewares.AuthAdminMiddleware(c.DeleteAdmin, c.Store)).Methods("DELETE")

	r.HandleFunc("/api/v1/admin/editors", middlewares.AuthAdminMiddleware(c.GetEditors, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/editors", middlewares.AuthAdminMiddleware(c.CreateEditor, c.Store)).Methods("POST")
	r.HandleFunc("/api/v1/admin/editors/{id}", middlewares.AuthAdminMiddleware(c.GetEditor, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/editors/{id}", middlewares.AuthAdminMiddleware(c.UpdateEditor, c.Store)).Methods("PATCH")
	r.HandleFunc("/api/v1/admin/editors/{id}", middlewares.AuthAdminMiddleware(c.DeleteEditor, c.Store)).Methods("DELETE")

	r.HandleFunc("/api/v1/admin/directions", middlewares.AuthAdminMiddleware(c.Directions, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/directions", middlewares.AuthAdminMiddleware(c.CreateDirection, c.Store)).Methods("POST")
	r.HandleFunc("/api/v1/admin/directions/{id}", middlewares.AuthAdminMiddleware(c.Direction, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/directions/{id}", middlewares.AuthAdminMiddleware(c.UpdateDirection, c.Store)).Methods("PATCH")
	r.HandleFunc("/api/v1/admin/directions/{id}", middlewares.AuthAdminMiddleware(c.DeleteDirection, c.Store)).Methods("DELETE")

	r.HandleFunc("/api/v1/admin/courses", middlewares.AuthAdminMiddleware(c.Courses, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/courses", middlewares.AuthAdminMiddleware(c.CreateCourse, c.Store)).Methods("POST")
	r.HandleFunc("/api/v1/admin/courses/{id}", middlewares.AuthAdminMiddleware(c.Course, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/courses/{id}", middlewares.AuthAdminMiddleware(c.UpdateCourse, c.Store)).Methods("PATCH")
	r.HandleFunc("/api/v1/admin/courses/{id}", middlewares.AuthAdminMiddleware(c.DeleteCourse, c.Store)).Methods("DELETE")

	r.HandleFunc("/api/v1/admin/students", middlewares.AuthAdminMiddleware(c.Students, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/students/{id}", middlewares.AuthAdminMiddleware(c.Student, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/admin/students/{id}", middlewares.AuthAdminMiddleware(c.UpdateStudent, c.Store)).Methods("PATCH")
	r.HandleFunc("/api/v1/admin/students/{id}", middlewares.AuthAdminMiddleware(c.DeleteStudent, c.Store)).Methods("DELETE")
}

func addEditorRoutes(r *mux.Router, c *controller.Controller) {
	r.HandleFunc("/api/v1/editor", middlewares.AuthEditorMiddleware(c.EditorHome, c.Store)).Methods("GET")

	r.HandleFunc("/api/v1/editor/login", c.EditorLogIn).Methods("POST")
	r.HandleFunc("/api/v1/editor/logout", c.EditorLogOut).Methods("POST")

	r.HandleFunc("/api/v1/editor/courses", middlewares.AuthEditorMiddleware(c.EditorCourses, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/editor/courses/{id}", middlewares.AuthEditorMiddleware(c.Course, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/editor/courses/{id}", middlewares.AuthEditorMiddleware(c.UpdateCourse, c.Store)).Methods("PATCH")

	r.HandleFunc("/api/v1/editor/students", middlewares.AuthEditorMiddleware(c.EditorStudents, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/editor/students/{id}", middlewares.AuthEditorMiddleware(c.Student, c.Store)).Methods("GET")
	r.HandleFunc("/api/v1/editor/students/{id}", middlewares.AuthEditorMiddleware(c.UpdateStudent, c.Store)).Methods("PATCH")
	r.HandleFunc("/api/v1/editor/students/{id}", middlewares.AuthEditorMiddleware(c.DeleteStudent, c.Store)).Methods("DELETE")
}

func addUserRoutes(r *mux.Router, c *controller.Controller) {
	r.HandleFunc("/api/v1/course/{id}", c.CoursePage).Methods("GET")
	r.HandleFunc("/api/v1/course/{id}", c.StudentCreate).Methods("POST")
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

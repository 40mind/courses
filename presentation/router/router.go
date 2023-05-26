package router

import (
	mw "courses/middlewares"
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
	r.HandleFunc("/api/v1/admin", mw.AuthAdminMiddleware(c.AdminHome, c.Store)).Methods("GET", "OPTIONS")

	r.HandleFunc("/api/v1/admin/login", c.AdminLogIn).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/admin/logout", c.AdminLogOut).Methods("POST", "OPTIONS")

	r.HandleFunc("/api/v1/admin/admins", mw.AuthAdminMiddleware(c.GetAdmins, c.Store)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/admin/admins", mw.AuthAdminMiddleware(c.CreateAdmin, c.Store)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/admin/admins/{id}", mw.AuthAdminMiddleware(c.GetAdmin, c.Store)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/admin/admins/{id}", mw.AuthAdminMiddleware(c.DeleteAdmin, c.Store)).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/v1/admin/editors", mw.AuthAdminMiddleware(c.GetEditors, c.Store)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/admin/editors", mw.AuthAdminMiddleware(c.CreateEditor, c.Store)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/admin/editors/{id}", mw.AuthAdminMiddleware(c.GetEditor, c.Store)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/admin/editors/{id}", mw.AuthAdminMiddleware(c.UpdateEditor, c.Store)).Methods("PATCH", "OPTIONS")
	r.HandleFunc("/api/v1/admin/editors/{id}", mw.AuthAdminMiddleware(c.DeleteEditor, c.Store)).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/v1/admin/directions", mw.AuthAdminMiddleware(c.Directions, c.Store)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/admin/directions", mw.AuthAdminMiddleware(c.CreateDirection, c.Store)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/admin/directions/{id}", mw.AuthAdminMiddleware(c.Direction, c.Store)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/admin/directions/{id}", mw.AuthAdminMiddleware(c.UpdateDirection, c.Store)).Methods("PATCH", "OPTIONS")
	r.HandleFunc("/api/v1/admin/directions/{id}", mw.AuthAdminMiddleware(c.DeleteDirection, c.Store)).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/v1/admin/courses", mw.AuthAdminMiddleware(c.Courses, c.Store)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/admin/courses", mw.AuthAdminMiddleware(c.CreateCourse, c.Store)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/admin/courses/{id}", mw.AuthAdminMiddleware(c.Course, c.Store)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/admin/courses/{id}", mw.AuthAdminMiddleware(c.UpdateCourse, c.Store)).Methods("PATCH", "OPTIONS")
	r.HandleFunc("/api/v1/admin/courses/{id}", mw.AuthAdminMiddleware(c.DeleteCourse, c.Store)).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/v1/admin/students", mw.AuthAdminMiddleware(c.Students, c.Store)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/admin/students/{id}", mw.AuthAdminMiddleware(c.Student, c.Store)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/admin/students/{id}", mw.AuthAdminMiddleware(c.UpdateStudent, c.Store)).Methods("PATCH", "OPTIONS")
	r.HandleFunc("/api/v1/admin/students/{id}", mw.AuthAdminMiddleware(c.DeleteStudent, c.Store)).Methods("DELETE", "OPTIONS")
}

func addEditorRoutes(r *mux.Router, c *controller.Controller) {
	r.HandleFunc("/api/v1/editor", mw.AuthEditorMiddleware(c.EditorHome, c.Store)).Methods("GET", "OPTIONS")

	r.HandleFunc("/api/v1/editor/login", c.EditorLogIn).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/editor/logout", c.EditorLogOut).Methods("POST", "OPTIONS")

	r.HandleFunc("/api/v1/editor/courses", mw.AuthEditorMiddleware(c.EditorCourses, c.Store)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/editor/courses/{id}", mw.AuthEditorMiddleware(c.Course, c.Store)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/editor/courses/{id}", mw.AuthEditorMiddleware(c.UpdateCourse, c.Store)).Methods("PATCH", "OPTIONS")

	r.HandleFunc("/api/v1/editor/students", mw.AuthEditorMiddleware(c.EditorStudents, c.Store)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/editor/students/{id}", mw.AuthEditorMiddleware(c.Student, c.Store)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/editor/students/{id}", mw.AuthEditorMiddleware(c.UpdateStudent, c.Store)).Methods("PATCH", "OPTIONS")
	r.HandleFunc("/api/v1/editor/students/{id}", mw.AuthEditorMiddleware(c.DeleteStudent, c.Store)).Methods("DELETE", "OPTIONS")
}

func addUserRoutes(r *mux.Router, c *controller.Controller) {
	r.HandleFunc("/api/v1/course/{id}", c.CoursePage).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/v1/course/{id}", c.StudentCreate).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/payment/create/{id}", c.CreatePayment).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/payment/confirm/{id}", c.ConfirmPayment).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1", c.HomePage).Methods("GET", "OPTIONS")
}

func addTechRoutes(r *mux.Router, c *controller.Controller) {
	r.HandleFunc("/api/v1/tech/info", c.TechInfo).Methods("GET")
}

func addFrontRoutes(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "./front/index.html")})
	r.HandleFunc("/course/{id}", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "./front/course_detail.html")})
	r.HandleFunc("/student/record/{id}", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "./front/create_student.html")})
	r.HandleFunc("/student/payment/{id}", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "./front/confirm_payment.html")})

	r.HandleFunc("/admin/login", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "./front/login.html")})
	r.HandleFunc("/admin/panel", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "./front/admin_panel.html")})

	r.HandleFunc("/editor/login", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "./front/editor_login.html")})
	r.HandleFunc("/editor/panel", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "./front/editor_panel.html")})

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./front/static"))))
}

package controller

import (
    "courses/domain/models"
    "encoding/json"
    "github.com/gorilla/sessions"
    "github.com/jimlawless/whereami"
    "gopkg.in/guregu/null.v4"
    "io"
    "log"
    "net/http"
    "strconv"
    "strings"
)

const sessionCookieName = "admin-session"

func (c *Controller) AdminHome(w http.ResponseWriter, r *http.Request) {
    directions, err := c.Service.GetDirections(r.Context(), "")
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    courses, err := c.Service.GetCourses(r.Context(), -1, "")
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    students, err := c.Service.GetStudents(r.Context(), "")
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    admins, err := c.Service.GetAdmins(r.Context(), "")
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    info := models.AdminInfo{
        Direction: directions,
        Course:    courses,
        Student:   students,
        Admins:    admins,
    }

    responseJson, err := json.Marshal(info)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, responseJson, nil, http.StatusOK)
}

func (c *Controller) AdminGetAdmins(w http.ResponseWriter, r *http.Request) {
    searchStr := r.URL.Query().Get("search")

    admins, err := c.Service.GetAdmins(r.Context(), searchStr)
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    responseJson, err := json.Marshal(admins)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, responseJson, nil, http.StatusOK)
}

func (c *Controller) AdminGetAdmin(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    admin, err := c.Service.GetAdmin(r.Context(), id)
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    nullAdmin := models.Admin{}
    if admin == nullAdmin {
        writeResponse(w, nil, nil, http.StatusNotFound)
        return
    }

    responseJson, err := json.Marshal(admin)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, responseJson, nil, http.StatusOK)
}

func (c *Controller) DeleteAdmin(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    err = c.Service.DeleteAdmin(r.Context(), id)
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, nil, nil, http.StatusOK)
}

func (c *Controller) AdminCreateAdmin(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    var admin models.Admin
    err = json.Unmarshal(body, &admin)
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    err, status := c.Service.CreateAdmin(r.Context(), admin)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    writeResponse(w, nil, nil, http.StatusCreated)
}

func (c *Controller) AdminLogIn(w http.ResponseWriter, r *http.Request) {
    var admin models.Admin
    err := json.NewDecoder(r.Body).Decode(&admin)
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    admin, err, status := c.Service.AdminLogIn(r.Context(), admin.Login, admin.Password)
    if err != nil || status != http.StatusOK {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, status)
        return
    }

    session, err := c.Store.Get(r, sessionCookieName)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    session.Values["login"] = admin.Login.String
    session.Values["authenticated"] = true
    session.Options = &sessions.Options{
        Path: "/",
        MaxAge:   86400,
    }
    err = session.Save(r, w)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, nil, nil, http.StatusOK)
}

func (c *Controller) AdminLogOut(w http.ResponseWriter, r *http.Request) {
    session, err := c.Store.Get(r, sessionCookieName)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    session.Values["authenticated"] = false
    session.Options = &sessions.Options{
        Path: "/",
        MaxAge:   -1,
    }
    err = session.Save(r, w)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, nil, nil, http.StatusOK)
}

func (c *Controller) AdminDirections(w http.ResponseWriter, r *http.Request) {
    searchStr := r.URL.Query().Get("search")

    directions, err := c.Service.GetDirections(r.Context(), searchStr)
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    responseJson, err := json.Marshal(directions)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, responseJson, nil, http.StatusOK)
}

func (c *Controller) AdminDirection(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    direction, err := c.Service.GetDirection(r.Context(), id)
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }
    emptyDirection := models.Direction{}
    if direction == emptyDirection {
        writeResponse(w, nil, nil, http.StatusNotFound)
        return
    }

    responseJson, err := json.Marshal(direction)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, responseJson, nil, http.StatusOK)
}

func (c *Controller) AdminCreateDirection(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    var direction models.Direction
    err = json.Unmarshal(body, &direction)
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    err, status := c.Service.CreateDirection(r.Context(), direction)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    writeResponse(w, nil, nil, http.StatusCreated)
}

func (c *Controller) AdminUpdateDirection(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    var direction models.Direction
    err = json.Unmarshal(body, &direction)
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }
    direction.Id = null.IntFrom(int64(id))

    err, status := c.Service.UpdateDirection(r.Context(), direction)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    writeResponse(w, nil, nil, http.StatusOK)
}

func (c *Controller) AdminDeleteDirection(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    err = c.Service.DeleteDirection(r.Context(), id)
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, nil, nil, http.StatusOK)
}

func (c *Controller) AdminCourses(w http.ResponseWriter, r *http.Request) {
    searchStr := r.URL.Query().Get("search")

    courses, err := c.Service.GetCourses(r.Context(), -1, searchStr)
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    responseJson, err := json.Marshal(courses)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, responseJson, nil, http.StatusOK)
}

func (c *Controller) AdminCourse(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    course, err := c.Service.GetCourse(r.Context(), id)
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }
    emptyCourse := models.Course{}
    if course == emptyCourse {
        writeResponse(w, nil, nil, http.StatusNotFound)
        return
    }

    responseJson, err := json.Marshal(course)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, responseJson, nil, http.StatusOK)
}

func (c *Controller) AdminCreateCourse(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    var course models.Course
    err = json.Unmarshal(body, &course)
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    err, status := c.Service.CreateCourse(r.Context(), course)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    writeResponse(w, nil, nil, http.StatusCreated)
}

func (c *Controller) AdminUpdateCourse(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    var course models.Course
    err = json.Unmarshal(body, &course)
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }
    course.Id = null.IntFrom(int64(id))

    err, status := c.Service.UpdateCourse(r.Context(), course)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    writeResponse(w, nil, nil, http.StatusOK)
}

func (c *Controller) AdminDeleteCourse(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    err = c.Service.DeleteCourse(r.Context(), id)
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, nil, nil, http.StatusOK)
}

func (c *Controller) AdminStudents(w http.ResponseWriter, r *http.Request) {
    searchStr := r.URL.Query().Get("search")

    students, err := c.Service.GetStudents(r.Context(), searchStr)
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    responseJson, err := json.Marshal(students)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, responseJson, nil, http.StatusOK)
}

func (c *Controller) AdminStudent(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    student, err := c.Service.GetStudent(r.Context(), id)
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }
    emptyStudent := models.Student{}
    if student == emptyStudent {
        writeResponse(w, nil, nil, http.StatusNotFound)
        return
    }

    responseJson, err := json.Marshal(student)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, responseJson, nil, http.StatusOK)
}

func (c *Controller) AdminCreateStudent(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    var student models.Student
    err = json.Unmarshal(body, &student)
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    _, err, status := c.Service.CreateStudent(r.Context(), student)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    writeResponse(w, nil, nil, http.StatusCreated)
}

func (c *Controller) AdminUpdateStudent(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    var student models.Student
    err = json.Unmarshal(body, &student)
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }
    student.Id = null.IntFrom(int64(id))

    err, status := c.Service.UpdateStudent(r.Context(), student)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    writeResponse(w, nil, nil, http.StatusOK)
}

func (c *Controller) AdminDeleteStudent(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    err = c.Service.DeleteStudent(r.Context(), id)
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, nil, nil, http.StatusOK)
}

package controller

import (
    "courses/domain/models"
    "courses/service"
    "encoding/json"
    "github.com/jimlawless/whereami"
    "log"
    "net/http"
    "strconv"
    "strings"
)

const controllerError = "controller error"

type Controller struct {
    Service *service.Service
}

func NewController(service *service.Service) *Controller {
    return &Controller{
        Service: service,
    }
}

func (c *Controller) HomePage(w http.ResponseWriter, r *http.Request) {
    courses, err := c.Service.GetCourses(r.Context())
    if err != nil {
        writeResponse(w, nil, err, 500)
        return
    }

    responseJson, err := json.Marshal(courses)
    if err != nil {
        log.Printf("%s: %s: %s", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, 500)
        return
    }

    writeResponse(w, responseJson, nil, 200)
}

func (c *Controller) CoursePage(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        writeResponse(w, nil, err, 400)
    }

    course, err := c.Service.GetCourse(r.Context(), id)
    if err != nil {
        writeResponse(w, nil, err, 500)
        return
    }

    responseJson, err := json.Marshal(course)
    if err != nil {
        log.Printf("%s: %s: %s", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, 500)
        return
    }

    writeResponse(w, responseJson, nil, 200)
}

func (c *Controller) CreateStudent(w http.ResponseWriter, r *http.Request) {
    err, status := c.Service.CreateStudent(r.Context(), r)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    writeResponse(w, nil, nil, status)
}

func (c *Controller) AdminHome(w http.ResponseWriter, r *http.Request) {
    directions, err := c.Service.GetDirections(r.Context())
    if err != nil {
        writeResponse(w, nil, err, 500)
        return
    }

    courses, err := c.Service.GetCourses(r.Context())
    if err != nil {
        writeResponse(w, nil, err, 500)
        return
    }

    students, err := c.Service.GetStudents(r.Context())
    if err != nil {
        writeResponse(w, nil, err, 500)
        return
    }

    info := models.AdminInfo{
        Directions: directions,
        Courses:    courses,
        Students:   students,
    }

    responseJson, err := json.Marshal(info)
    if err != nil {
        log.Printf("%s: %s: %s", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, 500)
        return
    }

    writeResponse(w, responseJson, nil, 200)
}

func (c *Controller) AdminUpdateCourse(w http.ResponseWriter, r *http.Request) {
    err, status := c.Service.UpdateCourse(r.Context(), r)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    writeResponse(w, nil, nil, 200)
}

func (c *Controller) AdminDeleteCourse(w http.ResponseWriter, r *http.Request) {
    err, status := c.Service.DeleteCourse(r.Context(), r)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    writeResponse(w, nil, nil, 200)
}

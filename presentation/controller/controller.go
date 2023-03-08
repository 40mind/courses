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
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }
    if courses == nil {
        writeResponse(w, nil, nil, http.StatusNoContent)
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

func (c *Controller) CoursePage(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        writeResponse(w, nil, err, http.StatusBadRequest)
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
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    courses, err := c.Service.GetCourses(r.Context())
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    students, err := c.Service.GetStudents(r.Context())
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    info := models.AdminInfo{
        Direction: directions,
        Course:    courses,
        Student:   students,
    }

    responseJson, err := json.Marshal(info)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, responseJson, nil, http.StatusOK)
}

func (c *Controller) AdminUpdateCourse(w http.ResponseWriter, r *http.Request) {
    err, status := c.Service.UpdateCourse(r.Context(), r)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    writeResponse(w, nil, nil, http.StatusOK)
}

func (c *Controller) AdminDeleteCourse(w http.ResponseWriter, r *http.Request) {
    err, status := c.Service.DeleteCourse(r.Context(), r)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    writeResponse(w, nil, nil, http.StatusOK)
}

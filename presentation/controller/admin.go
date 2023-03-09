package controller

import (
    "courses/domain/models"
    "encoding/json"
    "github.com/jimlawless/whereami"
    "gopkg.in/guregu/null.v4"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
    "strings"
)

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

func (c *Controller) AdminDirections(w http.ResponseWriter, r *http.Request) {
    directions, err := c.Service.GetDirections(r.Context())
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
    body, err := ioutil.ReadAll(r.Body)
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

    body, err := ioutil.ReadAll(r.Body)
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
    courses, err := c.Service.GetCourses(r.Context())
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

func (c *Controller) AdminStudents(w http.ResponseWriter, r *http.Request) {
    students, err := c.Service.GetStudents(r.Context())
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
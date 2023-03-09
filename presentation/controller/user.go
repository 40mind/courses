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

func (c *Controller) CreateStudent(w http.ResponseWriter, r *http.Request) {
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

    var student models.Student
    err = json.Unmarshal(body, &student)
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }
    student.CourseId = null.IntFrom(int64(id))

    err, status := c.Service.CreateStudent(r.Context(), student)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    writeResponse(w, nil, nil, http.StatusCreated)
}

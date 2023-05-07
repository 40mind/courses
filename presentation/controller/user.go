package controller

import (
    "courses/domain/models"
    "encoding/json"
    "github.com/jimlawless/whereami"
    "gopkg.in/guregu/null.v4"
    "io"
    "log"
    "net/http"
    "strconv"
    "strings"
)

func (c *Controller) HomePage(w http.ResponseWriter, r *http.Request) {
    dir := r.URL.Query().Get("direction")
    searchStr := r.URL.Query().Get("search")

    dirInt := -1
    var err error
    if dir != "" {
        dirInt, err = strconv.Atoi(dir)
        if err != nil {
            writeResponse(w, nil, err, http.StatusBadRequest)
            return
        }
    }

    courses, err := c.Service.GetCourses(r.Context(), dirInt, searchStr)
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    directions, err := c.Service.GetDirections(r.Context(), "")
    if err != nil {
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    var info struct{
        Directions  []models.Direction  `json:"directions"`
        Courses     []models.Course     `json:"courses"`
    }

    info.Directions = directions
    info.Courses = courses

    responseJson, err := json.Marshal(info)
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
        writeResponse(w, nil, nil, http.StatusNoContent)
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
    student.CourseId = null.IntFrom(int64(id))

    studentId, err, status := c.Service.CreateStudent(r.Context(), student)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    var responseBody struct{
        StudentId   int64 `json:"student_id"`
    }
    responseBody.StudentId = studentId.Int64
    responseJson, err := json.Marshal(responseBody)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, responseJson, nil, http.StatusCreated)
}

func (c *Controller) CreatePayment(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    paymentUrl, err, status := c.Service.CreatePayment(r.Context(), id, c.config.Server.Host)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    var response struct{
        ConfirmationUrl     string `json:"confirmation_url"`
    }
    response.ConfirmationUrl = paymentUrl
    responseJson, err := json.Marshal(response)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, responseJson, nil, http.StatusOK)
}

func (c *Controller) ConfirmPayment(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    err, status := c.Service.ConfirmPayment(r.Context(), id)
    if err != nil {
        writeResponse(w, nil, err, status)
        return
    }

    writeResponse(w, nil, nil, http.StatusOK)
}
package controller

import (
    "courses/domain/models"
    "encoding/json"
    "github.com/gorilla/sessions"
    "github.com/jimlawless/whereami"
    "log"
    "net/http"
    "strconv"
)

const editorSessionCookieName = "editor-session"

func (c *Controller) EditorHome(w http.ResponseWriter, r *http.Request) {
    writeResponse(w, nil, nil, http.StatusOK)
}

func (c *Controller) EditorLogIn(w http.ResponseWriter, r *http.Request) {
    var editor models.Editor
    err := json.NewDecoder(r.Body).Decode(&editor)
    if err != nil {
        log.Printf("%s: %s: %s\n", badRequest, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusBadRequest)
        return
    }

    err, status := c.Service.EditorLogIn(r.Context(), editor.Login, editor.Password)
    if err != nil || status != http.StatusOK {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, status)
        return
    }

    session, err := c.Store.Get(r, editorSessionCookieName)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    session.Values["role"] = "editor"
    session.Values["login"] = editor.Login.String
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

func (c *Controller) EditorLogOut(w http.ResponseWriter, r *http.Request) {
    session, err := c.Store.Get(r, editorSessionCookieName)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    session.Options = &sessions.Options{
        Path: "/",
        MaxAge: -1,
    }
    err = session.Save(r, w)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, nil, nil, http.StatusOK)
}

func (c *Controller) EditorCourses(w http.ResponseWriter, r *http.Request) {
    searchStr := r.URL.Query().Get("search")
    session, err := c.Store.Get(r, editorSessionCookieName)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }
    login, ok := session.Values["login"].(string)
    if !ok || login == "" {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    editor, err := c.Service.GetEditorByLogin(r.Context(), login)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    if editor.Courses == nil {
        writeResponse(w, nil, nil, http.StatusOK)
    }

    courses, err := c.Service.GetCourses(r.Context(), -1, searchStr, editor.Courses)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
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

func (c *Controller) EditorStudents(w http.ResponseWriter, r *http.Request) {
    searchStr := r.URL.Query().Get("search")
    course := r.URL.Query().Get("course")

    courseInt := -1
    var err error
    if course != "" {
        courseInt, err = strconv.Atoi(course)
        if err != nil {
            writeResponse(w, nil, err, http.StatusBadRequest)
            return
        }
    }

    session, err := c.Store.Get(r, editorSessionCookieName)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }
    login, ok := session.Values["login"].(string)
    if !ok || login == "" {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    editor, err := c.Service.GetEditorByLogin(r.Context(), login)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    if editor.Courses == nil {
        writeResponse(w, nil, nil, http.StatusOK)
    }

    students, err := c.Service.GetStudents(r.Context(), courseInt, searchStr, editor.Courses)
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

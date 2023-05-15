package middlewares

import (
    "github.com/gorilla/sessions"
    "log"
    "net/http"
)

const (
	adminSessionCookieName = "admin-session"
    editorSessionCookieName = "editor-session"
)

func AuthAdminMiddleware(handler http.HandlerFunc, store *sessions.CookieStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        session, err := store.Get(r, adminSessionCookieName)
        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        if role, ok := session.Values["role"].(string); !ok || role != "admin" {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        handler.ServeHTTP(w, r)
    }
}

func AuthEditorMiddleware(handler http.HandlerFunc, store *sessions.CookieStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        session, err := store.Get(r, editorSessionCookieName)
        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        if role, ok := session.Values["role"].(string); !ok || role != "editor" {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        handler.ServeHTTP(w, r)
    }
}

package middlewares

import (
    "github.com/gorilla/sessions"
    "log"
    "net/http"
)

const sessionCookieName = "admin-session"

func AuthMiddleware(handler http.HandlerFunc, store *sessions.CookieStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        session, err := store.Get(r, sessionCookieName)
        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        handler.ServeHTTP(w, r)
    }
}

package controller

import (
	"net/http"
)

func writeResponse(w http.ResponseWriter, body []byte, err error, status int) {
	w.WriteHeader(status)

	if err == nil {
		w.Write(body)
		return
	}

	if status / 100 == 5 {
		w.Write([]byte(`server error`))
		return
	}

	w.Write([]byte(err.Error()))
}

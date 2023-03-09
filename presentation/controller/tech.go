package controller

import (
    "encoding/json"
    "github.com/jimlawless/whereami"
    "log"
    "net/http"
)

func (c *Controller) TechInfo(w http.ResponseWriter, r *http.Request) {
    info, err := json.Marshal(c.config.Application)
    if err != nil {
        log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
        writeResponse(w, nil, err, http.StatusInternalServerError)
        return
    }

    writeResponse(w, info, nil, http.StatusOK)
}

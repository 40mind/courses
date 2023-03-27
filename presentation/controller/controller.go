package controller

import (
    "courses/domain/models"
    "courses/service"
    "github.com/gorilla/sessions"
)

const (
    controllerError = "controller error"
    badRequest = "bad request"
)

type Controller struct {
    Service *service.Service
    config  models.Config
    Store   *sessions.CookieStore
}

func NewController(service *service.Service, config models.Config) *Controller {
    store := sessions.NewCookieStore([]byte(config.Session.Key))
    return &Controller{
        Service: service,
        config:  config,
        Store:   store,
    }
}

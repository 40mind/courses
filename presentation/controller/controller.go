package controller

import (
    "courses/domain/models"
    "courses/service"
)

const (
    controllerError = "controller error"
    badRequest = "bad request"
)

type Controller struct {
    Service *service.Service
    config  models.Config
}

func NewController(service *service.Service, config models.Config) *Controller {
    return &Controller{
        Service: service,
        config:  config,
    }
}

package main

import (
	"courses/bootstrap"
	"courses/presentation/controller"
	"courses/presentation/router"
	"courses/service"
	"log"
	"net/http"
)

func main() {
	config := bootstrap.InitConfig()
	rep := bootstrap.InitRepository(&config.DB)
	svc := service.NewService(rep)
	con := controller.NewController(svc)
	r := router.NewRouter(con)
	log.Fatal(http.ListenAndServe(config.Server.Port, r))
}

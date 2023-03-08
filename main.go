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
	con := controller.NewController(svc, config)
	r := router.NewRouter(con)

	log.Printf("server started on port %s\n", config.Server.Port)
	log.Fatal(http.ListenAndServe(config.Server.Port, r))
}

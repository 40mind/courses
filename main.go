package main

import (
	"courses/bootstrap"
	"courses/presentation/controller"
	"courses/presentation/router"
	yookassaprovider "courses/providers/yookassa_provider"
	"courses/service"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	config := bootstrap.InitConfig()
	rep := bootstrap.InitRepository(config.DB)
	emailSender := bootstrap.InitEmailSender(config.Email)

	yookassaProv := yookassaprovider.NewProvider(config.YookassaProvider, config.YookassaAuth)
	svc := service.NewService(rep, emailSender, yookassaProv)
	con := controller.NewController(svc, config)
	r := router.NewRouter(con)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{config.Server.Host, config.Server.Host + config.Server.Port},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	log.Printf("server started on port %s\n", config.Server.Port)
	log.Fatal(http.ListenAndServe(config.Server.Port, c.Handler(r)))
}

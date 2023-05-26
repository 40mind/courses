package main

import (
	"context"
	"courses/bootstrap"
	"courses/domain/models"
	"courses/presentation/controller"
	"courses/presentation/router"
	yookassaprovider "courses/providers/yookassa_provider"
	"courses/service"
	"github.com/jimlawless/whereami"
	"github.com/rs/cors"
	"gopkg.in/guregu/null.v4"
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
	initAdmin(config, svc)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{config.Server.Host, config.Server.Host + config.Server.Port},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	log.Printf("server started on port %s\n", config.Server.Port)
	log.Fatal(http.ListenAndServe(config.Server.Port, c.Handler(r)))
}

func initAdmin(conf models.Config, svc *service.Service) {
	admins, err := svc.GetAdmins(context.Background(), "")
	if err != nil {
		log.Fatalf("init admin error: %s: %s", err.Error(), whereami.WhereAmI())
	}

	if admins != nil {
		return
	}

	err, _ = svc.CreateAdmin(context.Background(), models.Admin{
		Login:    null.StringFrom(conf.DefaultAdmin.Login),
		Password: null.StringFrom(conf.DefaultAdmin.Password),
	})
	if err != nil {
		log.Fatalf("init admin error: %s: %s", err.Error(), whereami.WhereAmI())
	}
}

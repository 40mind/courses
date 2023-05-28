package bootstrap

import (
	"courses/domain/models"
	"github.com/BurntSushi/toml"
	"github.com/jimlawless/whereami"
	"log"
	"os"
)

const configPath = "./configs/config.toml"

func InitConfig() models.Config {
	var config models.Config
	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		log.Fatalf("init config error: %s: %s", err.Error(), whereami.WhereAmI())
	}

	getEnv(&config)

	return config
}

func getEnv(config *models.Config) {
	var ok bool
	config.DB.User, ok = os.LookupEnv("DB_USER"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
	config.DB.Password, ok = os.LookupEnv("DB_PASSWORD"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
	config.DB.Name, ok = os.LookupEnv("DB_NAME"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
	config.Session.Key, ok = os.LookupEnv("SESSION_KEY"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
	config.Email.From, ok = os.LookupEnv("EMAIL_FROM"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
	config.Email.Password, ok = os.LookupEnv("EMAIL_PASSWORD"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
	config.YookassaAuth.Login, ok = os.LookupEnv("YOOKASSA_LOGIN"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
	config.YookassaAuth.Password, ok = os.LookupEnv("YOOKASSA_PASSWORD"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
	config.DefaultAdmin.Login, ok = os.LookupEnv("DEFAULT_ADMIN_LOGIN"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
	config.DefaultAdmin.Password, ok = os.LookupEnv("DEFAULT_ADMIN_PASSWORD"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
}

package bootstrap

import (
	"courses/domain/models"
	"github.com/BurntSushi/toml"
	"github.com/jimlawless/whereami"
	"log"
	"os"
	"path/filepath"
)

func InitConfig() models.Config {
	dev, ok := os.LookupEnv("DEVELOPMENT"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}

	configPath := ""
	if dev == "true" {
		configPath = filepath.Join("configs", "config.dev.toml")
	} else {
		configPath = filepath.Join("configs", "config.toml")
	}

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
	config.DB.User, ok = os.LookupEnv("DBUSER"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
	config.DB.Password, ok = os.LookupEnv("DBPASS"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
	config.DB.Host, ok = os.LookupEnv("DBHOST"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
	config.DB.Name, ok = os.LookupEnv("DBNAME"); if !ok {
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
	config.Server.Ip, ok = os.LookupEnv("APP_IP"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
	config.Server.Host, ok = os.LookupEnv("APP_HOST"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
	config.Server.Port, ok = os.LookupEnv("APP_PORT"); if !ok {
		log.Fatalf("init config error: %s", whereami.WhereAmI())
	}
}

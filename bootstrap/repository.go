package bootstrap

import (
	"courses/domain/models"
	"courses/domain/repository"
	"database/sql"
	"github.com/jimlawless/whereami"
	_ "github.com/lib/pq"
	"log"
)

func InitRepository(dbConfig models.DB) *repository.Repository {
	connectionString := "postgresql://" + dbConfig.User + ":" + dbConfig.Password + "@" + dbConfig.Host +
		dbConfig.Port + "/" + dbConfig.Name + dbConfig.ModeString

	db, err := sql.Open(dbConfig.Driver, connectionString)
	if err != nil {
		log.Fatalf("init repository error: %s: %s", err.Error(), whereami.WhereAmI())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("init repository error: %s: %s", err.Error(), whereami.WhereAmI())
	}

	return repository.New(db, dbConfig)
}

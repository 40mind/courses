package bootstrap

import (
	"courses/domain/models"
	"courses/domain/repository"
	"database/sql"
)

func InitRepository(dbConfig *models.DB) *repository.Repository {
	return repository.New(&sql.DB{}, dbConfig)
}

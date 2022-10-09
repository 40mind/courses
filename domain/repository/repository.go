package repository

import (
	"courses/domain/models"
	"database/sql"
)

type Repository struct {
	DB			*sql.DB
	DBConfig	*models.DB
}

func New(db *sql.DB, dbConfig *models.DB) *Repository {
	return &Repository{
		DB:       db,
		DBConfig: dbConfig,
	}
}

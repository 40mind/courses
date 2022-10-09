package service

import "courses/domain/repository"

type Service struct {
	Repository		*repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

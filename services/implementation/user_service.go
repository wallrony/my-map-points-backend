package services

import (
	"my_map_points/core/domain/models"
	repo_abs "my_map_points/data/repositories/abstraction"
)

type UserService struct {
	repository repo_abs.UserAbstractRepository
}

func (service UserService) FindAll() ([]models.User, models.ApiError) {
	var result []models.User

	result, err := service.repository.FindAll()

	return result, err
}

func NewUserService(repo repo_abs.UserAbstractRepository) *UserService {
	return &UserService{repository: repo}
}

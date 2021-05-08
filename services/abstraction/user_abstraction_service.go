package services

import "my_map_points/core/domain/models"

type UserAbstractService interface {
	FindAll() ([]models.User, models.ApiError)
}

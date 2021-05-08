package repository_abstraction

import "my_map_points/core/domain/models"

type UserAbstractRepository interface {
	FindAll() ([]models.User, models.ApiError)
}

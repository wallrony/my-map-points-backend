package repositories

import (
	abstraction "my_map_points/data/repositories/abstraction"

	"my_map_points/core/domain/models"
	"my_map_points/data/db"
)

type UserRepository struct {
	abstraction.UserAbstractRepository
}

func (repo *UserRepository) FindAll() ([]models.User, models.ApiError) {
	database, err := db.GetConnection()

	if err != nil {
		return nil, models.ApiError{Err: err.Error()}
	}

	defer database.Close()

	rows, err := database.Query("SELECT users.id, users.name, users.birthdate, users.email FROM users;")

	if err != nil {
		return nil, models.ApiError{Err: err.Error()}
	}

	list := []models.User{}

	for rows.Next() {
		var item models.User

		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Birthdate,
			&item.Email,
		)

		if err != nil {
			return nil, models.ApiError{Err: err.Error()}
		}

		list = append(list, item)
	}

	print(list == nil)

	return list, models.ApiError{}
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

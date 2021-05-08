package models

type User struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	Birthdate  string `db:"birthdate"`
	Email      string `db:"email"`
	PointCount int64  `db:"count(map_point)"`
}

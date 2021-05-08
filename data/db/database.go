package db

import (
	"fmt"
	"my_map_points/core/utils"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnection() (*sqlx.DB, error) {
	dbUser := utils.GetVar("DB_USER")
	dbPswd := utils.GetVar("DB_PSWD")
	dbName := utils.GetVar("DB_NAME")
	dbHost := utils.GetVar("DB_HOST")
	dbPort := utils.GetVar("DB_PORT")

	dataSource := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPswd, dbHost, dbPort, dbName)

	db, err := sqlx.Open("postgres", dataSource)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

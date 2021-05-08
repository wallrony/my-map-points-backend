package utils

import (
	"github.com/joho/godotenv"

	"os"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
}

func GetVar(key string) string {
	v := os.Getenv(key)

	return v
}

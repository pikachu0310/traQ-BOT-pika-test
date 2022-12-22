package util

import (
	"os"

	"github.com/joho/godotenv"
)

func GetToken() (token string, err error) {
	err = godotenv.Load(".env")
	token = os.Getenv("token")
	return token, err
}

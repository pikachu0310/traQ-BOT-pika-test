package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetToken() (token string) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error: tokenが読み込めなかった!: %v", err)
	}
	token = os.Getenv("token")
	return token
}

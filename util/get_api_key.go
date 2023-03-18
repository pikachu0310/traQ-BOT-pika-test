package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetApiKey() (apiKey string) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error: tokenが読み込めなかった!: %v", err)
	}
	apiKey = os.Getenv("apiKey")
	return apiKey
}

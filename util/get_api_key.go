package util

import (
	"os"
)

func GetApiKey() (apiKey string) {
	// 	err := godotenv.Load(".env")
	// 	if err != nil {
	// 		fmt.Printf("error: tokenが読み込めなかった!: %v", err)
	// 	}
	apiKey = os.Getenv("APIKEY")
	return apiKey
}

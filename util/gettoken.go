package util

import (
	"os"
)

func GetToken() (token string) {
	//err := godotenv.Load(".env")
	//if err != nil {
	//	fmt.Printf("error: tokenが読み込めなかった!: %v", err)
	//}
	token = os.Getenv("TOKEN")
	return token
}

package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("failed to get env variable: %v", err)
	}
	return os.Getenv(key)
}
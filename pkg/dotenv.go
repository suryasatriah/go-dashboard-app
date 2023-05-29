package pkg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file, error: %v", err)
	}
}

func GetDotEnvVariable(v string) string {
	loadDotEnv()

	variable := os.Getenv(v)
	if variable == "" {
		log.Fatalf("Environment variable is not set: %s", v)
	}

	return variable
}

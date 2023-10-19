package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// PORT returns the server listening port
	PORT = getEnv("PORT", "8081")
	// DB returns the name of the sqlite database
	DB = getEnv("DB", "gotodo.db")
	// TOKENKEY returns the jwt token secret
	TOKEN_KEY = getEnv("TOKEN_KEY", "")
	// TOKENEXP returns the jwt token expiration duration.
	// Should be time.ParseDuration string. Source: https://golang.org/pkg/time/#ParseDuration
	// default: 10h
	TOKENEXP = getEnv("TOKEN_EXP", "10h")
	// ALLOW_SITES returns the allowed sites for CORS
	ALLOW_SITES = getEnv("ALLOW_SITES", "")

	// MONGO_DB returns the mongodb database name
	MONGO_DB = getEnv("MONGO_DB", "")
)

func getEnv(name string, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}

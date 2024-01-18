package initializers

import (
	"log"
	"os"

	"github.com/Code-Mozart/minerva/minerva-backend/errors"
	"github.com/joho/godotenv"
)

type Environment struct {
	Port        string
	DatabaseURL string
}

func LoadEnvironment() (*Environment, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	databaseURL, err := getWithError("DATABASE_URL")
	if err != nil {
		return nil, err
	}

	return &Environment{
		Port:        getWithDefault("PORT", "8134"),
		DatabaseURL: databaseURL,
	}, nil
}

func getWithError(key string) (string, error) {
	value, present := os.LookupEnv(key)
	if !present {
		return "", &errors.MissingEnvironmentVariableError{
			Key: key,
		}
	}
	return value, nil
}

func getWithDefault(key string, defaultValue string) string {
	value, present := os.LookupEnv(key)
	if !present {
		log.Printf("Environment variable '%s' not found, using default value '%s'", key, defaultValue)
		return defaultValue
	}
	return value
}

package server

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const MAX_CONNECTION_ATTEMPTS = 5
const CONNECTION_ATTEMPT_DELAY = 5

func (server *Server) InitDatabase() (*gorm.DB, error) {
	connectString := server.Environment.DatabaseURL

	var lastError error
	for i := 0; i < MAX_CONNECTION_ATTEMPTS; i++ {
		database, err := server.tryOpenDatabase(connectString)
		lastError = err
		if err == nil {
			return database, nil
		}

		if i < MAX_CONNECTION_ATTEMPTS-1 {
			fmt.Printf("Failed to connect to database. Retrying in %d seconds...\n", CONNECTION_ATTEMPT_DELAY)
			time.Sleep(CONNECTION_ATTEMPT_DELAY * time.Second)
		}
	}

	return nil, lastError
}

func (server *Server) tryOpenDatabase(connectString string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(connectString), &gorm.Config{})
}

package config

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strconv"
)

func GetMongoURI() (*options.ClientOptions, string, error) {
	host := os.Getenv("MONGODB_HOST")
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	dbName := os.Getenv("MONGODB_DB_NAME")
	port, err := strconv.Atoi(os.Getenv("MONGODB_PORT"))

	if err != nil || len(host) == 0 || len(username) == 0 || len(password) == 0 || len(dbName) == 0 {
		fmt.Println("Missing MongoDB Configuration")
		return nil, "", errors.New("no MongoDB Configuration was provided")
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/", username, password, host, port)

	return options.Client().ApplyURI(uri), dbName, nil
}

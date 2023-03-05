package db

import (
	"errors"
	"github.com/dre-zouhair/mailer/internal/config"
	"github.com/go-redis/redis"
)

type Database struct {
	db *redis.Client
}

func Connect() (*Database, error) {
	dbConf := config.NewRedisConf()
	if dbConf == nil {
		return nil, errors.New("no DB conf was provided")
	}

	client := redis.NewClient(dbConf.GetRedisOptions())

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	} else {
		return &Database{
			client,
		}, nil
	}
}

func (db *Database) Close() error {
	return db.db.Close()
}

func (db *Database) GetDB() *redis.Client {
	return db.db
}

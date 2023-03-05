package model

import (
	"context"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

type Target struct {
	Entity
	Email string            `json:"email"`
	Model map[string]string `json:"model"`
}

type TargetRepository struct {
	*Repository[Target]
}

func NewTargetRepository(conn *redis.Client) *TargetRepository {
	return &TargetRepository{
		Repository: &Repository[Target]{
			conn:  conn,
			name:  "targets",
			entry: new(Target),
		},
	}
}

type TargetMongoRepository struct {
	*MongoRepository[Target]
}

func NewTargetMongoRepository(connection *mongo.Database) *TargetMongoRepository {
	return &TargetMongoRepository{
		&MongoRepository[Target]{
			"targets",
			connection,
			context.Background(),
		},
	}
}

package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Target struct {
	Email string            `json:"email"`
	Model map[string]string `json:"model"`
}

type TargetRepository struct {
	*Repository[Target]
}

func NewTargetMongoRepository(connection *mongo.Database) *TargetRepository {
	return &TargetRepository{
		&Repository[Target]{
			"targets",
			connection,
			context.Background(),
		},
	}
}

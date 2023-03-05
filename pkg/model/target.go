package model

import "go.mongodb.org/mongo-driver/mongo"

type Target struct {
	Email string            `json:"email"`
	Model map[string]string `json:"model"`
	Entity
}

type TargetRepository struct {
	*Repository[Target]
}

func NewTargetRepository(connection *mongo.Database) *TargetRepository {
	return &TargetRepository{
		&Repository[Target]{
			"targets",
			connection,
		},
	}
}

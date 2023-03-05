package model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Template struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Entity
}

type TemplateRepository struct {
	*Repository[Template]
}

func NewTemplateRepository(connection *mongo.Database) *TemplateRepository {
	return &TemplateRepository{
		&Repository[Template]{
			"targets",
			connection,
		},
	}
}

package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type Template struct {
	ID      primitive.ObjectID `json:"_id"`
	Name    string             `json:"name"`
	Subject string             `json:"subject"`
	Body    string             `json:"body"`
}

func (template *Template) Build(model map[string]string) (body string) {
	body = template.Body
	for key, value := range model {
		if len(value) == 0 {
			fmt.Println("no valid value for " + key + " in " + template.ID.Hex())
		}
		body = strings.ReplaceAll(body, key, value)
	}
	return body
}

type TemplateRepository struct {
	*Repository[Template]
	*MongoRepository[Template]
}

func NewTemplateRepository(conn *redis.Client) *TemplateRepository {
	return &TemplateRepository{
		Repository: &Repository[Template]{
			conn:  conn,
			name:  "templates",
			entry: new(Template),
		},
	}
}

type TemplateMongoRepository struct {
	*MongoRepository[Template]
}

func NewTemplateMongoRepository(connection *mongo.Database) *TemplateMongoRepository {
	return &TemplateMongoRepository{
		&MongoRepository[Template]{
			"targets",
			connection,
			context.Background(),
		},
	}
}

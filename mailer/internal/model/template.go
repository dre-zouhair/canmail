package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type Template struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type TemplateRepository struct {
	*Repository[Template]
}

func (template *Template) Build(model map[string]string) (body string) {
	body = template.Body
	for key, value := range model {
		if len(value) == 0 {
			fmt.Println("no valid value for " + key + " in " + template.Name)
		}
		body = strings.ReplaceAll(body, key, value)
	}
	return body
}

func NewTemplateMongoRepository(connection *mongo.Database) *TemplateRepository {
	return &TemplateRepository{
		&Repository[Template]{
			"templates",
			connection,
			context.Background(),
		},
	}
}

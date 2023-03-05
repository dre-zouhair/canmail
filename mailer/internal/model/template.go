package model

import (
	"fmt"
	"strings"
)

type Template struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

func (template *Template) Build(model map[string]string) {
	for key, value := range model {
		if value == "" {
			fmt.Println("no valid value for " + key + " in " + template.ID)
		}
		template.Body = strings.ReplaceAll(template.Body, key, value)
	}
}

type TemplateRepository struct {
	*Repository[Template]
}

func NewTemplateRepository() *TemplateRepository {
	return &TemplateRepository{
		Repository: &Repository[Template]{
			name:  "templates",
			entry: new(Template),
		},
	}
}

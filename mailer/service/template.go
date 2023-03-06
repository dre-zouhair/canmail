package service

import (
	"context"
	"github.com/dre-zouhair/mailer/internal/db"
	"github.com/dre-zouhair/mailer/internal/model"
)

type TemplateService struct {
	repo *model.TemplateMongoRepository
}

func NewTemplateService() *TemplateService {
	connection, dbName := db.GetMongoDBConnection(context.Background())

	templateRepository := model.NewTemplateMongoRepository(connection.Database(dbName))
	return &TemplateService{
		templateRepository,
	}
}

func (service *TemplateService) AddTemplate(t model.Template) int64 {
	return service.repo.Save([]model.Template{t})
}

func (service *TemplateService) UpdateTemplate(t model.Template) int64 {
	return service.repo.UpdateByID(t.ID.Hex(), t)
}

func (service *TemplateService) RemoveTemplate(t model.Template) int64 {
	return service.repo.DeleteOne(t.ID.Hex())
}

func (service *TemplateService) Get(id string) *model.Template {
	return service.repo.FindOne(id)
}

func (service *TemplateService) Retrieve(page, limit int64) []model.Template {
	return service.repo.Paginate(page, limit)
}

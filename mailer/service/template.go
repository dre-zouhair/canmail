package service

import (
	"github.com/dre-zouhair/mailer/internal/db"
	"github.com/dre-zouhair/mailer/internal/model"
)

type TemplateService struct {
	repo *model.TemplateRepository
}

func NewTemplateService() *TemplateService {
	connection, err := db.Connect()
	if err != nil {
		return nil
	}
	defer closeConnect(connection)
	templateRepository := model.NewTemplateRepository(connection.GetDB())
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

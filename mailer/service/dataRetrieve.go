package service

import (
	"context"
	"fmt"
	"github.com/dre-zouhair/mailer/internal/db"
	"github.com/dre-zouhair/mailer/internal/model"
)

func RetrieveBulkData(templateName string) (*model.Template, []model.Target) {
	return PaginateBulkData(templateName, 1, -1)
}

func PaginateBulkData(templateName string, page, limit int64) (*model.Template, []model.Target) {
	client, dbName := db.GetDBConnection(context.Background())
	defer client.Disconnect(context.Background())
	database := client.Database(dbName)
	templateRepository := model.NewTemplateMongoRepository(database)
	targetRepository := model.NewTargetMongoRepository(database)
	if limit == -1 {
		limit = targetRepository.Count()
	}
	return retrieveBulkData(templateName, templateRepository, targetRepository, page, limit)
}

func retrieveBulkData(templateName string, templateRepository *model.TemplateRepository, targetRepository *model.TargetRepository, page, limit int64) (*model.Template, []model.Target) {
	template := templateRepository.FindBy("name", templateName)
	if template == nil {
		fmt.Println("No template was found with the name " + templateName)
		return nil, nil
	}

	targets := targetRepository.Paginate(page, limit)
	if targets == nil || len(targets) == 0 {
		fmt.Println("No targets associated with the template " + templateName + " were found")
		return nil, nil
	}
	return &template[0], targets
}

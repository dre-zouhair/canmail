package service

import (
	"fmt"
	"github.com/dre-zouhair/mailer/internal/db"
	"github.com/dre-zouhair/mailer/internal/model"
)

func RetrieveBulkData(templateName string) (*model.Template, []model.Target) {

	connection, err := db.Connect()
	if err != nil {
		return nil, nil
	}
	defer closeConnect(connection)
	templateRepository := model.NewTemplateRepository(connection.GetDB())
	template := templateRepository.Get("id", templateName)

	targetRepository := model.NewTargetRepository(connection.GetDB())
	targets := targetRepository.GetAll()

	return template, targets

}

func closeConnect(connection *db.Database) {
	err := connection.Close()
	if err != nil {
		fmt.Println("unable to close redis connection")
		return
	} else {
		fmt.Println("Disconnected")
	}
}

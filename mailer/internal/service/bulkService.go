package service

import (
	"errors"
	"fmt"
	"github.com/dre-zouhair/mailer/internal/config"
	"github.com/dre-zouhair/mailer/internal/db"
	"github.com/dre-zouhair/mailer/internal/model"
)

func BulkMails(tempName string) []error {
	connection, err := db.Connect()
	if err != nil {
		return []error{
			errors.New("unable to init connection with redis db : " + err.Error()),
		}
	}

	mailServer, err := GetMailServer()
	if err != nil {
		return []error{
			errors.New("unable to init mail server : " + err.Error()),
		}
	}

	defer closeConnect(connection)
	templateRepository := model.NewTemplateRepository(connection.GetDB())
	template := templateRepository.Get(tempName, tempName)

	targetRepository := model.NewTargetRepository(connection.GetDB())
	targets := targetRepository.GetAll()
	return bulkMail(targets, template, mailServer)
}

func bulkMail(targets []model.Target, template *model.Template, mailServer *MailServer) []error {
	mails := make([]Mail, 0)
	for _, target := range targets {
		templateBody := template.Build(target.Model)
		mails = append(mails, Mail{
			To:      []string{target.Email},
			Subject: template.Subject,
			Body:    templateBody,
		})
	}

	errs := mailServer.SendMails(mails)

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func GetMailServer() (*MailServer, error) {
	mailConf := config.NewSMTPConf()

	if !mailConf.HasFromConf() {
		return nil, errors.New("no mail configuration was provided for from conf")
	}

	mailServer := NewMailServer("", mailConf.Host, mailConf.Domain, mailConf.GetFromConf().FromAddress, mailConf.GetFromConf().FromPassword)
	return mailServer, nil
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

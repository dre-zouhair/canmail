package service

import (
	"errors"
	"fmt"
	"github.com/dre-zouhair/mailer/internal/config"
	"github.com/dre-zouhair/mailer/internal/db"
	"github.com/dre-zouhair/mailer/internal/model"
)

func BulkMails(tempName string) {
	connection, err := db.Connect()
	if err != nil {
		fmt.Println("unable to init connection with redis db :", err)
		return
	}

	mailServer, err := GetMailServer()
	if err != nil {
		fmt.Println("unable to init mail server :", err)
		return
	}

	defer closeConnect(connection)
	templateRepository := model.NewTemplateRepository(connection.GetDB())
	template := templateRepository.Get(tempName, tempName)

	targetRepository := model.NewTargetRepository(connection.GetDB())
	targets := targetRepository.GetAll()

	for _, target := range targets {
		templateBody := template.Build(target.Model)
		err = mailServer.SendMail(Mail{
			To:      []string{target.Email},
			Subject: template.Subject,
			Body:    templateBody,
		})

		if err != nil {
			fmt.Println("unable to init connection with redis db :", err)
			return
		}
	}

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

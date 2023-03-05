package service

import (
	"errors"
	"github.com/dre-zouhair/mailer/internal/config"
	"github.com/dre-zouhair/mailer/internal/model"
)

func BulkMails(template *model.Template, targets []model.Target) []error {
	mailServer, err := GetMailServer()
	if err != nil {
		return []error{
			errors.New("unable to init mail server : " + err.Error()),
		}
	}
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

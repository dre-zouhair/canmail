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

func (server *MailServer) ConcurrentBulk(targets []model.Target, template *model.Template, success chan model.Target, fails chan model.Target) {
	if len(targets) < 50 {
		go server.chunk(targets, template, success, fails)
		return
	}

	i := 0
	for ; i < len(targets); i = i + 50 {
		go server.chunk(targets[i:i+50], template, success, fails)
	}

	go server.chunk(targets[i-50:], template, success, fails)

}

func (server *MailServer) chunk(targets []model.Target, template *model.Template, success chan model.Target, fails chan model.Target) {
	for _, target := range targets {
		templateBody := template.Build(target.Model)
		mail := Mail{
			To:      []string{target.Email},
			Subject: template.Subject,
			Body:    templateBody,
		}
		err := server.SendMail(mail)
		if err != nil {
			fails <- target
		} else {
			success <- target
		}
	}
}

package service

import (
	"errors"
	"fmt"
	"net/smtp"
)

type Mail struct {
	To         []string
	Subject    string
	Body       string
	Attachment []any
}

type MailServer struct {
	Address     string
	FromAddress string
	Auth        smtp.Auth
}

func NewMailServer(identity, host, address, fromAddress, fromPassword string) *MailServer {
	return &MailServer{
		Address:     address,
		FromAddress: fromPassword,
		Auth:        smtp.PlainAuth(identity, fromAddress, fromPassword, host),
	}
}

func (server *MailServer) SendMail(mail Mail) error {
	return smtp.SendMail(server.Address, server.Auth, server.FromAddress, mail.To, server.buildMailBody(mail))
}

func (server *MailServer) SendMails(mails []Mail) (errs []error) {
	for _, mail := range mails {
		err := server.SendMail(mail)
		if err != nil {
			errs = append(errs, errors.New("unable to send mail to : -> "+mail.To[0]+err.Error()))
		}
	}
	return errs
}

func (server *MailServer) buildMailBody(mail Mail) []byte {
	body := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n"+
		"%s", server.FromAddress, mail.To[0], mail.Subject, mail.Body)
	return []byte(body)
}

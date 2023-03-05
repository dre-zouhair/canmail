package config

import (
	"fmt"
	"os"
)

type SMTPConf struct {
	Domain     string // Address
	Host       string
	Port       int
	Username   string
	Password   string
	Encryption string
	from       *FromConf
}

type FromConf struct {
	FromAddress  string
	FromPassword string
	FromName     string
}

func (smtp *SMTPConf) hasFromCOnf() bool {
	return smtp.from != nil && smtp.from.FromAddress != "" && smtp.from.FromPassword != ""
}

func (smtp *SMTPConf) getFromCOnf() *FromConf {
	return smtp.from
}

func NewSMTPConf() *SMTPConf {
	domain := os.Getenv("SMTP_DOMAIN")
	if domain == "" {
		fmt.Println("SMTP_DOMAIN is not set")
		return nil
	}
	host := os.Getenv("SMTP_HOST")
	if host == "" {
		fmt.Println("SMTP_HOST is not set")
		return nil
	}

	username := os.Getenv("SMTP_USERNAME")
	if domain == "" {
		fmt.Println("SMTP_USERNAME is not set")
		return nil
	}

	password := os.Getenv("SMTP_PASSWORD")
	if host == "" {
		fmt.Println("SMTP_PASSWORD is not set")
		return nil
	}

	from := os.Getenv("SMTP_FROM")
	fromPass := os.Getenv("SMTP_FROM_PASSWORD")
	fromName := os.Getenv("SMTP_FROM_NAME")

	return &SMTPConf{
		Domain:   domain,
		Host:     host,
		Username: username,
		Password: password,
		from: &FromConf{
			FromAddress:  from,
			FromPassword: fromPass,
			FromName:     fromName,
		},
	}
}

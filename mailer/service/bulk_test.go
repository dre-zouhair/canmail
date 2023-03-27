package service

import (
	"fmt"
	"github.com/dre-zouhair/mailer/internal/model"
	"os"
	"strconv"
	"testing"
)

func Test_(t *testing.T) {

	os.Setenv("SMTP_DOMAIN", "localhost:1025")
	defer os.Getenv("SMTP_DOMAIN")
	os.Setenv("SMTP_HOST", "localhost")
	defer os.Getenv("SMTP_HOST")
	os.Setenv("SMTP_USERNAME", "smtp")
	defer os.Getenv("SMTP_USERNAME")
	os.Setenv("SMTP_PASSWORD", "password")
	defer os.Getenv("SMTP_PASSWORD")
	os.Setenv("SMTP_FROM", "from@mailer")
	defer os.Getenv("SMTP_FROM")
	os.Setenv("SMTP_FROM_PASSWORD", "password")
	defer os.Getenv("SMTP_FROM_PASSWORD")
	os.Setenv("SMTP_FROM_NAME", "from")
	defer os.Getenv("SMTP_FROM_NAME")

	type args struct {
		targets  []model.Target
		template *model.Template
	}

	targets := make([]model.Target, 0)
	for i := 0; i < 1034; i++ {
		targets = append(targets, model.Target{
			Email: "dre" + strconv.Itoa(i+1) + "@gm.com",
			Model: map[string]string{
				"name": "user " + strconv.Itoa(i+1),
			},
		})
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test cse 1",
			args{
				targets,
				&model.Template{
					Name:    "my test case template",
					Subject: "Test Mail",
					Body: "<h1>hi name !</h1>" +
						"This is an automatic mail",
				},
			},
		},
	}

	mailServer, err := GetMailServer()
	if err != nil {
		return
	}
	success := make(chan model.Target, 0)
	fails := make(chan model.Target, 0)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mailServer.ConcurrentBulk(tt.args.targets, tt.args.template, success, fails)
			for {
				select {
				case target := <-success:
					fmt.Println("success : ", target.Email)
				case target := <-fails:
					fmt.Println("fail : ", target.Email)
				}
			}
		})
	}
}

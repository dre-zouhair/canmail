package service

import (
	"net/smtp"
	"testing"
)

func TestMailServer_SendMail(t *testing.T) {
	type fields struct {
		address     string
		fromAddress string
		auth        smtp.Auth
	}
	type args struct {
		mail Mail
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			fields: fields{
				address:     "localhost:1025",
				fromAddress: "sender@example.com",
				auth:        smtp.PlainAuth("", "sender@test.com", "password", "localhost"),
			},
			args: args{
				Mail{
					To:      []string{"target@test.com"},
					Subject: "Test case 1",
					Body:    "<h1>HELLO</h1>",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := &MailServer{
				address:     tt.fields.address,
				fromAddress: tt.fields.fromAddress,
				auth:        tt.fields.auth,
			}
			if err := server.SendMail(tt.args.mail); (err != nil) != tt.wantErr {
				t.Errorf("SendMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

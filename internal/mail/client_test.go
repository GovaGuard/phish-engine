package mail

import (
	"os"
	"reflect"
	"testing"
)

func TestNewPlainMail(t *testing.T) {
	type args struct {
		sender  string
		subject string
		body    string
		to      []string
	}
	tests := []struct {
		name string
		args args
		want Mail
	}{
		{
			name: "SimplePlainMail", args: args{
				sender:  "mock@phish-engine.com",
				to:      []string{"target@phish-engine.com"},
				subject: "Gova Guard passt auf dich auf",
				body:    "Pass auf",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPlainMail(tt.args.sender, tt.args.subject, tt.args.body, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				if err := os.WriteFile(tt.name+".eml", got.Body, 0o755); err != nil {
					t.Log(err)
				}

				t.Errorf("NewPlainMail() = %v, want %v", got, tt.want)

			}
		})
	}
}

func TestSendMail(t *testing.T) {
	type args struct {
		mail Mail
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "SimplePlainMail", args: args{
				mail: Mail{
					Sender:  "mock@phish-engine.com",
					To:      []string{"target@phish-engine.com"},
					Subject: "Gova Guard passt auf dich auf",
					Body:    []byte("Pass auf"),
				},
			},
			wantErr: false,
		},
	}

	s := Sender{
		Sender:     "mock@phish-engine.com",
		User:       "mock",
		Password:   "mock",
		Host:       "127.0.0.1",
		SMTPServer: "127.0.0.1:2525",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewPlainMail(tt.args.mail.Sender, tt.args.mail.Subject, string(tt.args.mail.Body), tt.args.mail.To)

			if err := s.SendMail(m); (err != nil) != tt.wantErr {
				t.Errorf("NewPlainMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

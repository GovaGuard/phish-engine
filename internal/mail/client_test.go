package mail

import (
	"testing"
)

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

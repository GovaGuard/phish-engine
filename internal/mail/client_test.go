package mail

import "testing"

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
					Sender:  "nils.carstensen@govaguard.com",
					To:      []string{"me@nils-carstensen.de"},
					Subject: "Gova Guard passt auf dich auf",
					Body:    []byte("Pass auf"),
				},
			},
			wantErr: false,
		},
	}

	s := Sender{
		Sender:     "nils.carstensen@govaguard.com",
		User:       "nils.carstensen@govaguard.com",
		Password:   "456#SuppenKuecheGova",
		Host:       "w01ab24b.kasserver.com",
		SMTPServer: "w01ab24b.kasserver.com:25",
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

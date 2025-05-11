package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
)

const (
	boundary      = "\r\n--mail-boundary\r\n"
	boundaryIdent = "mail-boundary"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    []byte
}

type Sender struct {
	Sender   string
	User     string
	Password string
	// Host is only used for auth, usually it's the same as the SMTP server
	Host       string
	SMTPServer string
}

func (s *Sender) SendMail(m Mail) error {
	auth := smtp.PlainAuth("", s.User, s.Password, s.Host)

	// SendMail cannot perform DKIM signing. This is currently inherited by our E-Mail provided and technically not yet needed.
	if err := smtp.SendMail(s.SMTPServer, auth, m.Sender, m.To, m.Body); err != nil {
		return err
	}

	return nil
}

func NewPlainMail(sender, subject, body string, to []string) Mail {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("From: %s\r\n", sender))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(to, ";")))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString(fmt.Sprintf("Content-Type: text/plain; boundary=%s\n", boundaryIdent))

	buf.WriteString(boundary)
	buf.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	buf.WriteString(fmt.Sprintf("\r\n%s", body))
	buf.WriteString(boundary)

	return Mail{Sender: sender, To: to, Subject: subject, Body: buf.Bytes()}
}

func NewPlainMailWithAttachement(sender, subject, body string, to []string, file Attachement) Mail {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("From: %s\r\n", sender))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(to, ";")))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundaryIdent))

	buf.WriteString(boundary)
	buf.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	buf.WriteString(fmt.Sprintf("\r\n%s", body))

	buf.WriteString(boundary)
	buf.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	buf.WriteString("Content-Transfer-Encoding: base64\r\n")
	buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\r\n", file.Name))
	buf.WriteString(fmt.Sprintf("Content-ID: %s\r\n\r\n", file.Name))
	buf.Write(file.Base64Content)
	buf.WriteString(boundary)

	buf.WriteString("--")

	return Mail{Sender: sender, To: to, Subject: subject, Body: buf.Bytes()}
}

package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"path/filepath"
)

type MailService struct {
	smtpHost    string
	smtpPort    string
	from        string
	username    string
	password    string
	templateDir string
}

func NewMailService(
	host,
	port,
	from,
	username,
	password,
	tplDir string,
) MailServiceInterface {
	return &MailService{
		smtpHost:    host,
		smtpPort:    port,
		from:        from,
		username:    username,
		password:    password,
		templateDir: tplDir,
	}
}

func (s *MailService) Send(to, subject, tplFile string, data interface{}) error {
	tplPath := filepath.Join(s.templateDir, tplFile)
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := tpl.Execute(&body, data); err != nil {
		return err
	}

	message := []byte(
		"From: " + s.from + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0;\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\n\n" +
			body.String() + "\r\n")

	auth := smtp.PlainAuth("", s.username, s.password, s.smtpHost)
	serverAddress := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)

	return smtp.SendMail(serverAddress, auth, s.from, []string{to}, message)
}

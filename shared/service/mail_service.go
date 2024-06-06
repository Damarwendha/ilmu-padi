package service

import (
	"net/smtp"

	"github.com/kelompok-2/ilmu-padi/config"
)

type mailService struct {
	cfg config.SmtpConfig
}

type IMailService interface {
	SendMail(subject, html string, to []string) error
}

func (m *mailService) SendMail(subject, html string, to []string) error {
	auth := smtp.PlainAuth(
		"",
		m.cfg.EmailName,
		m.cfg.EmailAppPswd,
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + html

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		m.cfg.EmailName,
		to,
		[]byte(msg),
	)
	if err != nil {
		return err
	}

	return nil
}

func NewMailService(config config.SmtpConfig) IMailService {
	return &mailService{
		cfg: config,
	}
}

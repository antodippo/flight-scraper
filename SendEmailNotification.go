package main

import (
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load(".env.dist")
	}
	LogFatalAndExitIfNotNull(err)
}

type Server interface {
	SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) (err error)
	PlainAuth(identity, username, password, host string) (auth smtp.Auth)
}

type SMTPServer struct{}

func (server *SMTPServer) PlainAuth(identity, username, password, host string) (auth smtp.Auth) {
	return smtp.PlainAuth(identity, username, password, host)
}

func (server *SMTPServer) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) (err error) {
	return smtp.SendMail(addr, a, from, to, msg)
}

// SendEmailNotification sends an HTML email notification
func SendEmailNotification(server Server, to []string, body []byte, searchInput SearchInput) {
	from := os.Getenv("SMTP_FROM")
	auth := server.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("SMTP_PWD"), os.Getenv("SMTP_HOST"))
	msg := "From: " + from + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: Flight results " + searchInput.Departure + "-" + searchInput.Arrival + " on " + searchInput.Date + " \n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		string(body)
	err := server.SendMail(
		os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
		auth,
		from,
		to,
		[]byte(msg),
	)
	LogFatalAndExitIfNotNull(err)
}

package main

import (
	"net/smtp"
)

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
func SendEmailNotification(server Server, to []string, body []byte) {
	from := "flightScraper@antodippo.com"
	msg := "From: " + from + "\n" +
		"To: " + to[0] + "\n" +
		"Subject: Flight results\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		string(body)
	auth := server.PlainAuth("", "antodippo", "antomailer1", "smtp.sendgrid.net")
	err := server.SendMail(
		"smtp.sendgrid.net:587",
		auth,
		from,
		to,
		[]byte(msg),
	)
	LogFatalAndExitIfNotNull(err)
}

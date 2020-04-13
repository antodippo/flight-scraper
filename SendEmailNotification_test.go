package main

import (
	"net/smtp"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SpySMTPServer struct {
	To        []string
	Body      []byte
	AuthCalls int
	SendCalls int
}

func (server *SpySMTPServer) PlainAuth(identity, username, password, host string) (auth smtp.Auth) {
	server.AuthCalls++

	return
}

func (server *SpySMTPServer) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) (err error) {
	server.SendCalls++
	server.To = to
	server.Body = msg

	return
}

func TestSendEmailNotification(t *testing.T) {
	t.Run("It parses and returns flights results", func(t *testing.T) {
		server := &SpySMTPServer{}
		SendEmailNotification(server, []string{"test@test.com"}, []byte("TEST"))
		assert.Equal(t, 1, server.AuthCalls)
		assert.Equal(t, 1, server.SendCalls)
		assert.Equal(t, []string{"test@test.com"}, server.To)
		assert.Equal(t, []byte("TEST"), server.Body)
	})
}

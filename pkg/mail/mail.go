package mail

import (
	"strings"

	"gopkg.in/gomail.v2"
)

type Options struct {
	MailHost string
	MailPort int
	MailUser string
	MailPass string
	MailTo   string
	Subject  string
	Body     string
}

func Send(o *Options) error {
	m := gomail.NewMessage()

	m.SetHeader("From", o.MailUser)

	mailArrTo := strings.Split(o.MailTo, ",")
	m.SetHeader("To", mailArrTo...)

	m.SetHeader("Subject", o.Subject)

	m.SetBody("text/html", o.Body)

	d := gomail.NewDialer(o.MailHost, o.MailPort, o.MailUser, o.MailPass)
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d.DialAndSend(m)
}

package mailer

import (
	"gopkg.in/gomail.v2"
)

func Activationmail(email string, tokenn string) {

	m := gomail.NewMessage()
	m.SetHeader("From", "martinhercka1@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Hello")
	m.SetBody("text/html", "localhost:8080/activate/"+tokenn)

	d := gomail.NewDialer("smtp.gmail.com", 587, "martinhercka1", "shadowman1")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}

package mailer

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func Activationmail(email string) {
	fmt.Println("Hi")
	m := gomail.NewMessage()
	m.SetHeader("From", "martinhercka1@gmail.com")
	m.SetHeader("To", email)

	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello")

	d := gomail.NewDialer("smtp.gmail.com", 587, "mailSenderNick", "mailSenderPassword")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}

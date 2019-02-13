package mailer

import (
	"fmt"

	str "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
	gomail "gopkg.in/gomail.v2"
)

//Activationmail --
func Activationmail(email string, tokenn string, mailer str.Mail) {

	m := gomail.NewMessage()
	m.SetHeader("From", "martinhercka1@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Hello")
	m.SetBody("text/html", mailer.Host+":"+mailer.Port+"/auth/activate?token="+tokenn)

	if mailer.Username == "" || mailer.Password == "" {
		fmt.Println(mailer.Host + ":" + mailer.Port + "/auth/activate?token=" + tokenn)
		return
	}
	d := gomail.NewDialer("smtp.gmail.com", 587, mailer.Username, mailer.Password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}

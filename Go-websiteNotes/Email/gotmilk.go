package main

import (
	"fmt"
	"crypto/tls"
	gomail "gopkg.in/mail.v2"
)

func main() {

	Msg := gomail.NewMessage()

	// email Sender

	Msg.SetHeader("From", "From@gmail.com")

	// email Receiver

	Msg.SetHeader("To", "Receiver@outlook.com")

	// email subject

	Msg.SetHeader("Subject", "TestSubject1")

	// email body

	Msg.SetHeader("text/plain", "Body of email")

	// SMTP server settings 

	D := gomail.NewDialer("smtp.gmail.com", 587, "From@gmail.coom", "FromE-password")

	// if SSL/TSL crt is not valid on server.

	D.TLSConfig = &tls.Config{InsecureSkipVerify: true}


	// Send email

	if err := D.DialAndSend(Msg); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return

}


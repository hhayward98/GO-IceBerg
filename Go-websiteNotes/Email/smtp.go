package main

import (
	"fmt"
	"net/smtp"

)

func main() {

	SenderEmail := "example@gamil.com"
	Senderpassword := "Password for email"

	
	ReciverEmail := []string {"exampleR@outlook.com"}

	smtpHost := "smtp.gamil.com"
	smtpPort := "587"

	// Message
	Message := []byte("This is a Test message")

	// AUTH 
	auth := smtp.PlainAuth("", SenderEmail, Senderpassword, smtpHost)



	// Send 
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, SenderEmail, ReciverEmail, Message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully sent eamil!!!")
}
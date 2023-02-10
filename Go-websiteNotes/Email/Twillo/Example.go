package main

import (

	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {

	Source := mail.NewEmail("SourceEmail", "Source@gmail.com")

	Msg_Subject := "Email Subject Here"

	Destination := mail.NewEmail("Desination User", "Desination@gmail.com")

	PlainTextContent := "This is the main Message of the Email!"

	HTMLContent := "<strong>This is the main Message of the Email!</strong>"

	message := mail.NewSingleEmail(Source, Msg_Subject, Destination, PlainTextContent, HTMLContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

}



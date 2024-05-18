package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
)

func main() {
	from := mail.NewEmail("Example User", "tiar@wacaku.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "mawahyu@wacaku.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

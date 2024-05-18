package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
)

func main() {
	from := mail.NewEmail("Wacaku", "tiar@wacaku.com")
	subject := "Request Withdrawal"
	to := mail.NewEmail("Wahyu", "rupi@wacaku.com")
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

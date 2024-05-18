package main

import (
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	from := mail.NewEmail("Example User", "tiar@wacaku.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "mawahyu@wacaku.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	_ = mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	//response, err := client.Send(message)
	//if err != nil {
	//	log.Println(err)
	//} else {
	//	fmt.Println(response.StatusCode)
	//	fmt.Println(response.Body)
	//	fmt.Println(response.Headers)
	//}
}

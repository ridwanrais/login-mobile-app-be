package tools

// Import the Mailjet wrapper
import (
	"fmt"
	"log"
	"os"

	mailjet "github.com/mailjet/mailjet-apiv3-go/v4"
)

type SendEmailParams struct {
	RecipientEmail string
	RecipientName  string
	Subject        string
	HtmlContent    string
}

func SendEmail(params SendEmailParams) error {
	mailjetPublicKey := os.Getenv("MAILJET_API_KEY")
	mailjetPrivateKey := os.Getenv("MAILJET_SECRET_KEY")
	senderEmail := os.Getenv("MAILJET_SENDER_EMAIL")

	mailjetClient := mailjet.NewMailjetClient(mailjetPublicKey, mailjetPrivateKey)

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: senderEmail,
				Name:  "Ridwan Skyshi",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: params.RecipientEmail,
					Name:  params.RecipientName,
				},
			},
			Subject:  params.Subject,
			HTMLPart: params.HtmlContent,
			// TextPart: "Dear passenger 1, welcome to Mailjet! May the delivery force be with you!",
			// HTMLPart: "<h3>Dear passenger 1, welcome to <a href=\"https://www.mailjet.com/\">Mailjet</a>!</h3><br />May the delivery force be with you!",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	fmt.Printf("Email Data: %+v\n", res)

	return nil
}

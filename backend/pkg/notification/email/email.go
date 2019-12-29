package email

import (
	"fmt"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

type MailJetSender struct {
	cl        *mailjet.Client
	fromEmail string
	fromName  string
}

func NewMailJetSender(apiKey, secretKey, fromEmail, fromName string) *MailJetSender {
	return &MailJetSender{
		cl:        mailjet.NewMailjetClient(apiKey, secretKey),
		fromEmail: fromEmail,
		fromName:  fromEmail,
	}
}

func (s *MailJetSender) Send(toEmail, toName, subject, body string) error {
	msg := mailjet.InfoMessagesV31{
		From: &mailjet.RecipientV31{
			Email: s.fromEmail,
			Name:  s.fromName,
		},
		To: &mailjet.RecipientsV31{mailjet.RecipientV31{
			Email: toEmail,
			Name:  toName,
		}},
		Subject:  subject,
		TextPart: body,
	}
	_, err := s.cl.SendMailV31(&mailjet.MessagesV31{
		Info: []mailjet.InfoMessagesV31{
			msg,
		},
	})
	return fmt.Errorf(`failed to send email: [email: %s, err: %w]`, toEmail, err)
}

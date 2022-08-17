package mail

import (
	"encoding/base64"
	"errors"
	"log"

	"google.golang.org/api/gmail/v1"
)

func SendEmailOAUTH2(to string, data interface{}, template string) (bool, error) {
	OAuthGmailService()
	emailBody, err := ParseTemplate(template, data)
	if err != nil {
		log.Println(err, template)
		return false, errors.New("unable to parse email template")
	}

	var message gmail.Message

	emailTo := "To: " + to + "\r\n"
	subject := "Subject: " + "work" + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + emailBody)

	message.Raw = base64.URLEncoding.EncodeToString(msg)

	// Send the message
	_, err = GmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

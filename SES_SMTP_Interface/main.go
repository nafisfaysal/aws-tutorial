package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

const (
	// Replace SmtpUser with your Amazon SES SMTP user name.
	SmtpUser = ""

	// Replace SmtpPass with your Amazon SES SMTP password.
	SmtpPass = ""

	// If you're using Amazon SES in an AWS Region other than US West (Oregon),
	// replace email-smtp.us-west-2.amazonaws.com with the Amazon SES SMTP
	// endpoint in the appropriate region.
	Host = "email-smtp.us-west-2.amazonaws.com"
	Port = 587
)

func main() {
	SendEmail("jon@example.com", "Jon", "recipient@ex.com", "Hello AWS SES", "<p>535 Authentication Credentials Invalid</p>", "You did a great job. Congratulation", "aws,ses")

}

func SendEmail(sender, senderName, recipient, subject, htmlBody, textBody, tags string) {
	m := gomail.NewMessage()

	// Set the main email part to use HTML.
	m.SetBody("text/html", htmlBody)

	// Set the alternative part to plain text.
	m.AddAlternative("text/plain", textBody)

	// Construct the message headers, including a Configuration Set and a Tag.
	m.SetHeaders(map[string][]string{
		"From":               {m.FormatAddress(sender, senderName)},
		"To":                 {recipient},
		"Subject":            {subject},
		"X-SES-MESSAGE-TAGS": {tags},
	})

	// Send the email.
	d := gomail.NewPlainDialer(Host, Port, SmtpUser, SmtpPass)

	// Display an error message if something goes wrong; otherwise,
	// display a message confirming that the message was sent.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email sent!")
	}

}

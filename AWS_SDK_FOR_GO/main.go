package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	// Replace AccessKeyID with your AccessKeyID key.
	AccessKeyID = "23423randomstuff"

	// Replace AccessKeyID with your AccessKeyID key.
	SecretAccessKey = "23423randomstuff"

	// Replace sender@example.com with your "From" address.
	// This address must be verified with Amazon SES.
	Sender = "sender@example.com"

	// Replace us-west-2 with the AWS Region you're using for Amazon SES.
	AwsRegion = "us-west-2"

	// The character encoding for the email.
	CharSet = "UTF-8"
)

func main() {
	err := SendEmail("ex@ex.aws", "Hello AWS SES", "<p>535 Authentication Credentials Invalid</p>")
	if err != nil {
		fmt.Println("AWS SES SendEmail ERROR: ", err)
	}
}

// SendEmail delivery an email utilizing the AWS SES service
func SendEmail(recipient, subject, htmlBody string) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AwsRegion),
		Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
	},
	)

	svc := ses.New(sess)
	payload := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(htmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(htmlBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(Sender),
	}

	res, err := svc.SendEmail(payload)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				log.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				log.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				log.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				log.Println(aerr.Error())
			}
		} else {
			log.Println(err.Error())
		}
		return nil
	}

	fmt.Println("Email Sent!")
	fmt.Println(res)

	return nil
}

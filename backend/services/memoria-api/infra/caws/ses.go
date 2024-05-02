package caws

import (
	"context"
	"fmt"

	"memoria-api/domain/interfaces"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

const (
	// The character encoding for the email.
	CharSet = "UTF-8"
)

type sesMailer struct {
	client *sesv2.Client
}

func NewSESMailer(cfg aws.Config) (mailer interfaces.Mailer, err error) {
	client := sesv2.NewFromConfig(cfg)
	mailer = &sesMailer{client: client}
	return
}

func (m *sesMailer) Send(ctx context.Context, dto interfaces.MailerSendDTO) (err error) {
	input := &sesv2.SendEmailInput{
		FromEmailAddress: &dto.From,
		Destination: &types.Destination{
			ToAddresses: dto.To,
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Data: &dto.Body,
					},
					Html: &types.Content{
						Data: &dto.Body,
					},
				},
				Subject: &types.Content{
					Data: &dto.Subject,
				},
			},
		},
	}

	res, err := m.client.SendEmail(ctx, input)
	if err != nil {
		return
	}

	fmt.Println(res.MessageId)
	fmt.Println("SESMailer Send success")
	return
}

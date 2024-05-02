package interfaces

import "context"

type Mailer interface {
	Send(ctx context.Context, dto MailerSendDTO) error
}

type MailerSendDTO struct {
	From    string
	To      []string
	Subject string
	Body    string
}

package notifications

import (
	"log"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"


	"github.com/pshebel/partiburo/mail/env"
)

func Publish(party_id int, subject string, message string) error {
	ctx := context.Background()
	client := GetSesClient()


	for _, email := range(contacts) {
		 input := &sesv2.SendEmailInput{
			Destination: &types.Destination{
				ToAddresses: []string{
					email,
				},
			},
			Content: &types.EmailContent{
				Simple: &types.Message {
					Body: &types.Body{
						Text: &types.Content{
							Charset: aws.String("UTF-8"),
							Data:    aws.String(message),
						},
					},
					Subject: &types.Content{
						Charset: aws.String("UTF-8"),
						Data:    aws.String(subject),
					},
				},
			},
			FromEmailAddress: aws.String(env.AwsSender),
		}

		// Send the email
		_, err := client.SendEmail(ctx, input)
		if err != nil {
			return err
		}

	}
	return nil
}
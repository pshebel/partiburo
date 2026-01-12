package notifications

import (
	"log"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"

	"github.com/pshebel/partiburo/backend/env"
)

func PublishEmail(email, subject, body string) error {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(env.AwsRegion))
	if err != nil {
		log.Fatal(err)
		return err
	}
	client := sesv2.NewFromConfig(cfg)

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
						Data:    aws.String(body),
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
	_, err = client.SendEmail(ctx, input)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
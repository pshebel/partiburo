package main

import (
	"log"
    "fmt"
	"database/sql"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"

	"github.com/pshebel/partiburo/mail/database"
	"github.com/pshebel/partiburo/mail/env"
	"github.com/pshebel/partiburo/mail/utils"

)

type Reminder struct {
	PartyId	string
	Guest	string
	Party	string
	Email	string
}

func sendBulk(recipients []Reminder) {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(env.AwsRegion))
	if err != nil {
		log.Println(err)
		return
	}
	client := sesv2.NewFromConfig(cfg)

	for _, r := range(recipients) {
		subject := "You have a party coming up!"
		message := fmt.Sprintf("Hey %s,\nYour %s party is coming up in 3 days.\n\nTo unsubscribe from reminders for this party, visit https://partiburo.com/unsubscribe/%s\nIf you want to unsubscribe from all future emails from partiburo, visit https://partiburo/unsubscribeAll/%s", r.Guest, r.Party, r.Email, r.Email)
		input := &sesv2.SendEmailInput{
			Destination: &types.Destination{
				ToAddresses: []string{
					r.Email,
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
			log.Println(err)
			return
		}

	}
}

func main() {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return 
	}
	party_id := 0

	query := `SELECT 
		p.id, p.title, g.name, g.email 
		FROM guests as g 
		LEFT JOIN party as p on p.id = g.party_id
		WHERE NOT EXISTS (
			SELECT 1
			FROM blacklist b
			WHERE b.email = g.email
		)
		AND EXISTS (
			SELECT 1 
			FROM whitelist w 
			WHERE w.email = g.email 
			AND w.confirmed = true
    	)
	`
	rows, err := db.Query(query, party_id)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	recipients := []Reminder{}
	for rows.Next() {
		var id string
		var title string
		var name string
		var emailNull sql.NullString
		err := rows.Scan(&id, &title, &name, &emailNull)
		if err != nil {
			fmt.Println(err)
			return
		}
		var email string
		if emailNull.Valid{
			email = emailNull.String
		} else {
			email = ""
		}
		
		if utils.IsValidEmail(email) {
			r := Reminder{
				PartyId:	id,
				Guest:		name,
				Party: 		title,
				Email: 		email,
			}
			recipients = append(recipients, r)
		}
	} 
	fmt.Println(recipients)
	sendBulk(recipients)
}
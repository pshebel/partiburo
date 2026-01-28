package main

import (
	"log"
    "fmt"
	"database/sql"
	"time"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"

	"github.com/pshebel/partiburo/mail/database"
	"github.com/pshebel/partiburo/mail/env"
	// "github.com/pshebel/partiburo/mail/utils"

)

const ANNOUNCEMENT_SUBJECT = "A party you subscribe to has a new announcement"
const REMINDER_SUBJECT = "You have a party coming up!"

const LAYOUT = "2006-01-02 15:04"

type Reminder struct {
	Subject	string
	Body	string
	Email	string
}

type Party struct {
	id string
	title string
	date string
	time string
	dayOf bool
	dayBefore bool
	weekBefore bool
	announcements bool
}

func send(email, subject, body string) error {

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(env.AwsRegion))
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return err
	}

	return nil
}

// func sendBulk(recipients []Reminder) {
// 	ctx := context.Background()
// 	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(env.AwsRegion))
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	client := sesv2.NewFromConfig(cfg)

// 	for _, r := range(recipients) {
// 		subject := "You have a party coming up!"
// 		message := fmt.Sprintf("Hey %s,\nYour %s party is coming up in 3 days.\n\nTo unsubscribe from reminders for this party, visit https://partiburo.com/unsubscribe/%s\nIf you want to unsubscribe from all future emails from partiburo, visit https://partiburo/unsubscribeAll/%s", r.Guest, r.Party, r.Email, r.Email)
// 		input := &sesv2.SendEmailInput{
// 			Destination: &types.Destination{
// 				ToAddresses: []string{
// 					r.Email,
// 				},
// 			},
// 			Content: &types.EmailContent{
// 				Simple: &types.Message {
// 					Body: &types.Body{
// 						Text: &types.Content{
// 							Charset: aws.String("UTF-8"),
// 							Data:    aws.String(message),
// 						},
// 					},
// 					Subject: &types.Content{
// 						Charset: aws.String("UTF-8"),
// 						Data:    aws.String(subject),
// 					},
// 				},
// 			},
// 			FromEmailAddress: aws.String(env.AwsSender),
// 		}

// 		// Send the email
// 		_, err := client.SendEmail(ctx, input)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}

// 	}
// }

func main() {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return 
	}


	query := `
		SELECT 
			p.id, 
			p.title,
			p.date,
			p.time,
			r.day_of,
			r.day_before,
			r.week_before,
			r.announcements
		FROM party as p 
		LEFT JOIN reminders as r ON r.party_id = p.id
		WHERE date>date('now')`
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	defer rows.Close()
	parties := []Party{}
	for rows.Next() {
		p := Party{}

		err := rows.Scan(&p.id, &p.title, &p.date, &p.time, &p.dayOf, &p.dayBefore, &p.weekBefore, &p.announcements) 
		if err != nil {
			log.Println(err)
			return
		}

		parties = append(parties, p)
	}

	for _, p := range(parties) {
		subject := ""
		body := ""

		sent := false
		now := time.Now()
		datetime := p.date + " " + p.time
		layout := "2006-01-02 15:04"
		deadline, err := time.ParseInLocation(layout, datetime, time.Local)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}

		isSameDay := now.Year() == deadline.Year() && now.Month() == deadline.Month() && now.Day() == deadline.Day()

		tomorrow := now.AddDate(0, 0, 1)
		isDayBefore := tomorrow.Year() == deadline.Year() && tomorrow.Month() == deadline.Month() && tomorrow.Day() == deadline.Day()

		nextWeek := now.AddDate(0, 0, 7)
		isWeekBefore := nextWeek.Year() == deadline.Year() && 
						nextWeek.Month() == deadline.Month() && 
						nextWeek.Day() == deadline.Day()

		if p.dayOf && now.Before(deadline) && isSameDay {
			sent = true
			subject = REMINDER_SUBJECT
			body = fmt.Sprintf("Your %s party is coming up today at %s.", p.title, p.time)
		} else if (p.dayBefore && isDayBefore) {
			sent = true
			subject = REMINDER_SUBJECT
			body = fmt.Sprintf("Your %s party is tomorrow.", p.title)
		} else if (p.weekBefore && isWeekBefore) {
			sent = true
			subject = REMINDER_SUBJECT
			body = fmt.Sprintf("Your %s party is next week.", p.title)
		}
 
		fmt.Println(p.announcements, sent)
		if (p.announcements && !sent) {
			query := `
				SELECT header, body 
				FROM announcements 
				WHERE party_id = ? 
				AND date(created_at) = date('now')
				ORDER BY created_at DESC 
				LIMIT 1
			`

			var header, message string
			row := db.QueryRow(query, p.id)
			err := row.Scan(&header, &message)
			if err != nil &&  err != sql.ErrNoRows {
				log.Println(err)
				return
			}

			if err == sql.ErrNoRows {
				continue
			}
			fmt.Println(header, message)

			subject = ANNOUNCEMENT_SUBJECT
			body = fmt.Sprintf("%s\n\n %s", header, message)
		}

		if subject != ""  && body != "" {
			query := `
				SELECT e.email 
				FROM guests as g 
				LEFT JOIN email as e ON e.id=g.email_id
				WHERE g.party_id=?
				AND NOT EXISTS (
					SELECT 1 
					FROM blacklist AS b 
					WHERE b.email_id = g.email_id
				)
			`
			rows, err := db.Query(query, p.id)
			if err != nil {
				log.Println(err)
				return
			}

			defer rows.Close()
			// emails := []string{}
			for rows.Next() {
				var email string
				err := rows.Scan(&email)
				if err != nil {
					log.Println(err)
					return
				}
				send(email, subject, body)
			}
		}
	}


}
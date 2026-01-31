package main

import (
	"context"
	"database/sql"
	"flag" // 1. Added flag package
	"fmt"
	"time"
	"os"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"

	"github.com/pshebel/partiburo/mail/database"
	"github.com/pshebel/partiburo/mail/env"
)

const ANNOUNCEMENT_SUBJECT = "A party you subscribe to has a new announcement"
const REMINDER_SUBJECT = "You have a party coming up!"
const LAYOUT = "2006-01-02 15:04"

type Party struct {
	id            string
	title         string
	date          string
	time          string
	dayOf         bool
	dayBefore     bool
	weekBefore    bool
	announcements bool
}

func send(email, subject, body string) error {
	slog.Info("sending", email, subject, body)
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(env.AwsRegion))
	if err != nil {
		slog.Error("server failed to start", "error", err)
		return err
	}
	client := sesv2.NewFromConfig(cfg)

	input := &sesv2.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{email},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
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

	_, err = client.SendEmail(ctx, input)
	if err != nil {
		slog.Error("server failed to start", "error", err)
		return err
	}

	return nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true, // This adds the file and line number!
	}))
	slog.SetDefault(logger)
	slog.Info("server starting", "version", 0)

	// 2. Define and parse the dry-run flag
	dryRun := flag.Bool("dry-run", false, "print emails to console instead of sending")
	flag.Parse()

	db, err := database.GetDB()
	if err != nil {
		slog.Error("server failed to start", "error", err)
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
        WHERE date>=date('now')`
	rows, err := db.Query(query)
	if err != nil {
		slog.Error("server failed to start", "error", err)
		return
	}
	defer rows.Close()

	parties := []Party{}
	for rows.Next() {
		p := Party{}
		err := rows.Scan(&p.id, &p.title, &p.date, &p.time, &p.dayOf, &p.dayBefore, &p.weekBefore, &p.announcements)
		if err != nil {
			slog.Error("server failed to start", "error", err)
			return
		}
		parties = append(parties, p)
	}

	for _, p := range parties {
		slog.Info("checking party ", p.title)
		subject := ""
		body := ""
		sent := false
		now := time.Now()
		datetime := p.date + " " + p.time
		deadline, err := time.ParseInLocation(LAYOUT, datetime, time.Local)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}

		isSameDay := now.Year() == deadline.Year() && now.Month() == deadline.Month() && now.Day() == deadline.Day()
		tomorrow := now.AddDate(0, 0, 1)
		isDayBefore := tomorrow.Year() == deadline.Year() && tomorrow.Month() == deadline.Month() && tomorrow.Day() == deadline.Day()
		nextWeek := now.AddDate(0, 0, 7)
		isWeekBefore := nextWeek.Year() == deadline.Year() && nextWeek.Month() == deadline.Month() && nextWeek.Day() == deadline.Day()

		if p.dayOf && now.Before(deadline) && isSameDay {
			sent = true
			subject = REMINDER_SUBJECT
			body = fmt.Sprintf("Your %s party is coming up today at %s.", p.title, p.time)
		} else if p.dayBefore && isDayBefore {
			sent = true
			subject = REMINDER_SUBJECT
			body = fmt.Sprintf("Your %s party is tomorrow.", p.title)
		} else if p.weekBefore && isWeekBefore {
			sent = true
			subject = REMINDER_SUBJECT
			body = fmt.Sprintf("Your %s party is next week.", p.title)
		}

		if p.announcements && !sent {
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
			if err != nil && err != sql.ErrNoRows {
				slog.Error("server failed to start", "error", err)
				return
			}
			if err == sql.ErrNoRows {
				continue
			}
			subject = ANNOUNCEMENT_SUBJECT
			body = fmt.Sprintf("%s\n\n %s", header, message)
		}

		if subject != "" && body != "" {
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
				slog.Error("server failed to start", "error", err)
				return
			}
			defer rows.Close()

			for rows.Next() {
				var email string
				var e sql.NullString
				err := rows.Scan(&e)
				if err != nil {
					slog.Error("server failed to start", "error", err)
					return
				}
				if e.Valid{
					email = e.String
				} else {
					continue
				}

				// 3. Conditional logic for Dry Run
				if *dryRun {
					fmt.Printf("[DRY RUN] To: %s | Subject: %s\nBody: %s\n---\n", email, subject, body)
				} else {
					send(email, subject, body)
				}
			}
		}
	}
}
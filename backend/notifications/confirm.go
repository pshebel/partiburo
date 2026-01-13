package notifications

import (
	"fmt"
	"log"
	"database/sql"

	"github.com/pshebel/partiburo/backend/utils"
	"github.com/pshebel/partiburo/backend/database"
)

func ConfirmEmail(email string) (bool, error) {
	log.Println("Confirm email")
	// check for valid email
	if email == "" || !utils.IsValidEmail(email) {
		return false, nil
	}

	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return false, nil
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return false, err
	}
	// check if email is blacklisted
	var count int64
	blacklist := `SELECT COUNT(*) FROM blacklist WHERE email=?`
	row := db.QueryRow(blacklist, email)
	err = row.Scan(&count)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if count > 0 {
		return false, nil
	}
	// check if email is already confirmed
	var confirmed bool
	whitelist := `SELECT confirmed FROM whitelist WHERE email=?`
	row = db.QueryRow(whitelist, email)
	err = row.Scan(&confirmed)
	if err != nil && err != sql.ErrNoRows  {
		log.Println(err)
		return false, err
	}
	if err == sql.ErrNoRows {
		passcode := utils.RandomString()
		subject := "Confirm your email with Partiburo"
		message := fmt.Sprintf("To confirm your email, click this link https://partiburo.com/confirm/%s/%s\n\nIf you were not expecting this email, you do not need to take any action", email, passcode)
		query := `INSERT INTO notifications (email, summary) VALUES (?, ?)`
		_, err := db.Exec(query, email, "Confirmation Email")
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return false, nil
		}


		err = PublishEmail(email, subject, message)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return false, nil
		}
		query = `INSERT INTO whitelist (email, passcode) VALUES (?, ?)`
		_, err = db.Exec(query, email, passcode)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return false, nil
		}
		return false, nil
	}
	if !confirmed {
		return false, nil
	}
	return true, nil
}
package notifications

import (
	"fmt"
	"log"
	"database/sql"

	"github.com/pshebel/partiburo/backend/utils"
	"github.com/pshebel/partiburo/backend/database"
)

func newEmail(tx *sql.Tx, email string) error {
	fmt.Println("new email")
	code := utils.RandomString()
	passcode := utils.RandomString()
	subject := "Confirm your email with Partiburo"
	message := fmt.Sprintf("To confirm your email, click this link https://partiburo.com/confirm/%s/%s\n\nIf you were not expecting this email, you do not need to take any action", code, passcode)

	query := `INSERT INTO email (code, email) VALUES (?, ?)`
	res, err := tx.Exec(query, code, email)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	query = `INSERT INTO whitelist (email_id, passcode) VALUES (?, ?)`
	_, err = tx.Exec(query, id, passcode)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	query = `INSERT INTO notifications (email_id, summary) VALUES (?, ?)`
	_, err = tx.Exec(query, id, "Confirmation Email")
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	err = PublishEmail(email, subject, message)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	return nil
}

func ConfirmEmail(email string) (bool, error) {
	log.Println("Confirm email")
	// check for valid email
	if email == "" || !utils.IsValidEmail(email) {
		fmt.Println("invalid email")
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
	defer tx.Commit()

	email_id := 0
	query := `SELECT id FROM email WHERE email=?`
	row := tx.QueryRow(query, email)
	err = row.Scan(&email_id)
	if err != nil && err != sql.ErrNoRows  {
		log.Println(err)
		return false, err
	}
	if err == sql.ErrNoRows {
		err := newEmail(tx, email)
		if err != nil {
			log.Println(err)
			return false, err
		}
		return false, nil
	}

	// check if email is blacklisted
	var count int64
	blacklist := `SELECT COUNT(*) FROM blacklist WHERE email_id=?`
	row = db.QueryRow(blacklist, email_id)
	err = row.Scan(&count)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if count > 0 {
		log.Println("email blacklisted")
		return false, nil
	}

	// check f email is confirmed
	var confirmed bool
	query = `SELECT confirmed FROM whitelist WHERE email_id=?`
	row = tx.QueryRow(query, email_id)
	err = row.Scan(&confirmed)
	if err != nil && err != sql.ErrNoRows  {
		log.Println(err)
		return false, err
	}
	if err == sql.ErrNoRows {
		err := newEmail(tx, email)
		if err != nil {
			log.Println(err)
			return false, err
		}
		return false, nil
	}

	if !confirmed {
		fmt.Println("email not confirmed")
		return false, nil
	}
	

	
	return true, nil
}
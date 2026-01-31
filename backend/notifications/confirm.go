package notifications

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/pshebel/partiburo/backend/database"
	"github.com/pshebel/partiburo/backend/utils"
)

func newEmail(tx *sql.Tx, email string) error {
	slog.Info("creating new email record", "email", email)

	code := utils.RandomString()
	passcode := utils.RandomString()
	subject := "Confirm your email with Partiburo"
	message := fmt.Sprintf("To confirm your email, click this link https://partiburo.com/confirm/%s/%s\n\nIf you were not expecting this email, you do not need to take any action", code, passcode)

	query := `INSERT INTO email (code, email) VALUES (?, ?)`
	res, err := tx.Exec(query, code, email)
	if err != nil {
		tx.Rollback()
		slog.Error("failed to insert email", "email", email, "error", err)
		return fmt.Errorf("newEmail insertion: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		slog.Error("failed to get last insert id", "email", email, "error", err)
		return fmt.Errorf("newEmail lastID: %w", err)
	}

	query = `INSERT INTO whitelist (email_id, passcode) VALUES (?, ?)`
	_, err = tx.Exec(query, id, passcode)
	if err != nil {
		tx.Rollback()
		slog.Error("failed to insert whitelist", "email_id", id, "error", err)
		return fmt.Errorf("newEmail whitelist: %w", err)
	}

	query = `INSERT INTO notifications (email_id, summary) VALUES (?, ?)`
	_, err = tx.Exec(query, id, "Confirmation Email")
	if err != nil {
		tx.Rollback()
		slog.Error("failed to insert notification", "email_id", id, "error", err)
		return fmt.Errorf("newEmail notification: %w", err)
	}

	err = PublishEmail(email, subject, message)
	if err != nil {
		tx.Rollback()
		slog.Error("failed to publish email", "email", email, "error", err)
		return fmt.Errorf("newEmail publish: %w", err)
	}

	return nil
}

func ConfirmEmail(email string) (bool, error) {
	slog.Info("processing email confirmation", "email", email)
	
	email = utils.SanitizeEmail(email)
	if email == "" || !utils.IsValidEmail(email) {
		slog.Warn("invalid email format provided", "email", email)
		return false, nil
	}

	db, err := database.GetDB()
	if err != nil {
		slog.Error("database connection failed", "error", err)
		return false, fmt.Errorf("ConfirmEmail db access: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		slog.Error("failed to begin transaction", "error", err)
		return false, fmt.Errorf("ConfirmEmail tx begin: %w", err)
	}
	defer tx.Commit()

	var email_id int
	query := `SELECT id FROM email WHERE email=?`
	row := tx.QueryRow(query, email)
	err = row.Scan(&email_id)

	if err != nil && err != sql.ErrNoRows {
		slog.Error("query error checking email existence", "email", email, "error", err)
		return false, fmt.Errorf("ConfirmEmail existence check: %w", err)
	}

	if err == sql.ErrNoRows {
		if err := newEmail(tx, email); err != nil {
			return false, err // Error is already logged and wrapped in newEmail
		}
		return false, nil
	}

	// check if email is blacklisted
	var count int64
	blacklist := `SELECT COUNT(*) FROM blacklist WHERE email_id=?`
	err = db.QueryRow(blacklist, email_id).Scan(&count)
	if err != nil {
		slog.Error("failed to check blacklist", "email_id", email_id, "error", err)
		return false, fmt.Errorf("ConfirmEmail blacklist check: %w", err)
	}
	if count > 0 {
		slog.Warn("attempted access from blacklisted email", "email_id", email_id)
		return false, nil
	}

	// check if email is confirmed
	var confirmed bool
	query = `SELECT confirmed FROM whitelist WHERE email_id=?`
	err = tx.QueryRow(query, email_id).Scan(&confirmed)
	if err != nil && err != sql.ErrNoRows {
		slog.Error("failed to check whitelist confirmation", "email_id", email_id, "error", err)
		return false, fmt.Errorf("ConfirmEmail confirmed check: %w", err)
	}
	
	if err == sql.ErrNoRows {
		if err := newEmail(tx, email); err != nil {
			return false, err
		}
		return false, nil
	}

	if !confirmed {
		slog.Debug("email verification pending", "email_id", email_id)
		return false, nil
	}

	return true, nil
}
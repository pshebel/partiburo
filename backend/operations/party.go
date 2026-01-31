package operations

import (
	"database/sql"
	"fmt"
	"log/slog" // Replaced "log"
	"strings"

	"github.com/pshebel/partiburo/backend/database"
	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/notifications"
	"github.com/pshebel/partiburo/backend/utils"
)

func GetTitles(req models.TitlesRequest) (models.TitlesResponse, *models.Response) {
	slog.Info("fetching titles", "count", len(req.Codes))
	resp := models.TitlesResponse{}

	if len(req.Codes) == 0 {
		slog.Warn("no codes submitted for titles")
		return resp, &models.Response{Code: 404, Message: "no titles"}
	}

	db, err := database.GetDB()
	if err != nil {
		slog.Error("database connection failed", "error", err)
		return resp, &models.Response{Code: 500, Message: "Service Error"}
	}

	placeholders := make([]string, len(req.Codes))
	args := make([]interface{}, len(req.Codes))
	for i, v := range req.Codes {
		placeholders[i] = "?"
		args[i] = v
	}

	query := fmt.Sprintf(
		"SELECT user_code, title FROM party WHERE user_code IN (%s);",
		strings.Join(placeholders, ","),
	)

	rows, err := db.Query(query, args...)
	if err != nil {
		slog.Error("query failed", "query", query, "error", err)
		return resp, &models.Response{Code: 500, Message: "Service Error"}
	}
	defer rows.Close()

	results := make(map[string]string)
	for rows.Next() {
		var userCode, title string
		if err := rows.Scan(&userCode, &title); err != nil {
			slog.Error("row scan failed", "error", err)
			return resp, &models.Response{Code: 500, Message: "Service Error"}
		}
		results[userCode] = title
	}

	resp.Titles = results
	return resp, nil
}

func GetParty(code string) (models.Party, *models.Response) {
	slog.Info("fetching party details", "code", code)
	party := models.Party{}
	party_id := 0

	db, err := database.GetDB()
	if err != nil {
		slog.Error("database connection failed", "error", err)
		return party, &models.Response{Code: 500, Message: "Service Error"}
	}

	partyQuery := `SELECT id, date, time, address, title, description FROM party WHERE user_code = ? OR admin_code= ?`
	err = db.QueryRow(partyQuery, code, code).Scan(&party_id, &party.Date, &party.Time, &party.Address, &party.Title, &party.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("party not found", "code", code)
			return party, &models.Response{Code: 404, Message: "Could not find party"}
		}
		slog.Error("party fetch failed", "code", code, "error", err)
		return party, &models.Response{Code: 500, Message: "Service Error"}
	}

	// Reminders
	query := `SELECT day_of, day_before, week_before, announcements from reminders where party_id=?`
	var dayOf, dayBefore, weekBefore, announcements bool
	err = db.QueryRow(query, party_id).Scan(&dayOf, &dayBefore, &weekBefore, &announcements)
	if err != nil && err != sql.ErrNoRows {
		slog.Error("reminders fetch failed", "party_id", party_id, "error", err)
		return party, &models.Response{Code: 500, Message: "Service Error"}
	}
	
	if dayOf { party.Reminders = append(party.Reminders, "day_of") }
	if dayBefore { party.Reminders = append(party.Reminders, "day_before") }
	if weekBefore { party.Reminders = append(party.Reminders, "week_before") }
	if announcements { party.Reminders = append(party.Reminders, "announcements") }

	// Announcements
	query = `SELECT id, header, body, created_at FROM announcements WHERE party_id=?`
	rows, err := db.Query(query, party_id)
	if err != nil {
		slog.Error("announcements query failed", "party_id", party_id, "error", err)
		return party, &models.Response{Code: 500, Message: "service error"}
	}
	defer rows.Close()

	party.Announcements = []models.Announcement{}
	for rows.Next() {
		var a models.Announcement
		if err := rows.Scan(&a.ID, &a.Header, &a.Body, &a.CreatedAt); err != nil {
			slog.Error("announcement scan failed", "party_id", party_id, "error", err)
			return party, &models.Response{Code: 500, Message: "service error"}
		}
		party.Announcements = append(party.Announcements, a)
	}

	return party, nil
}

func CreateParty(req models.PartyRequest) (models.PartyResponse, *models.Response) {
	slog.Info("creating new party", "title", req.Title, "email", req.AdminEmail)
	resp := models.PartyResponse{}

	email := utils.SanitizeEmail(req.AdminEmail)
	if email == "" || !utils.IsValidEmail(email) {
		slog.Warn("invalid email format provided", "email", email)
		return resp,  &models.Response{Code: 400, Message: "Invalid email"}
	}

	confirmed, err := notifications.ConfirmEmail(email)
	if err != nil {
		slog.Error("email confirmation check failed", "email", email, "error", err)
		return resp, &models.Response{Code: 500, Message: "Service Error"}
	}

	db, err := database.GetDB()
	if err != nil {
		slog.Error("database connection failed", "error", err)
		return resp, &models.Response{Code: 500, Message: "Service Error"}
	}

	userCode := utils.RandomString()
	adminCode := utils.RandomString()

	tx, err := db.Begin()
	if err != nil {
		slog.Error("failed to begin transaction", "error", err)
		return resp, &models.Response{Code: 500, Message: "Service Error"}
	}
	defer tx.Rollback() // Safe to defer rollback; it does nothing if tx is committed

	var email_id int
	err = tx.QueryRow("SELECT id FROM email WHERE email=?", email).Scan(&email_id)
	if err != nil {
		slog.Error("failed to find email id", "email", email, "error", err)
		return resp, &models.Response{Code: 500, Message: "Service Error"}
	}

	res, err := tx.Exec("INSERT INTO party (admin_email_id, user_code, admin_code, date, time, address, title, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", 
		email_id, userCode, adminCode, req.Date, req.Time, req.Address, req.Title, req.Description)
	if err != nil {
		slog.Error("failed to insert party", "email_id", email_id, "error", err)
		return resp, &models.Response{Code: 500, Message: "Service Error"}
	}

	party_id, _ := res.LastInsertId()

	var ann, dOf, dBef, wBef bool
	for _, r := range req.Reminders {
		switch r {
		case "day_of": dOf = true
		case "day_before": dBef = true
		case "week_before": wBef = true
		case "announcements": ann = true
		}
	}

	_, err = tx.Exec("INSERT INTO reminders (party_id, day_of, day_before, week_before, announcements) VALUES (?, ?, ?, ?, ?)", 
		party_id, dOf, dBef, wBef, ann)
	if err != nil {
		slog.Error("failed to insert reminders", "party_id", party_id, "error", err)
		return resp, &models.Response{Code: 500, Message: "Service Error"}
	}

	subject := fmt.Sprintf("Here are your Partiburo links for %s", req.Title)
	message := fmt.Sprintf("Guest Link: https://partiburo.com/%s\nAdmin Link: https://partiburo.com/admin/%s\n", userCode, adminCode)
	
	if confirmed {
		if err := notifications.PublishEmail(email, subject, message); err != nil {
			slog.Error("failed to publish email", "email", email, "error", err)
			return resp, &models.Response{Code: 500, Message: "Service Error"}
		}
	} else {
		_, err = tx.Exec("INSERT INTO queue (email_id, subject, body, sent, retry) VALUES (?, ?, ?, ?, ?)", 
			email_id, subject, message, false, "daily")
		if err != nil {
			slog.Error("failed to queue email", "email_id", email_id, "error", err)
			return resp, &models.Response{Code: 500, Message: "Service Error"}
		}
	}

	if err := tx.Commit(); err != nil {
		slog.Error("transaction commit failed", "error", err)
		return resp, &models.Response{Code: 500, Message: "Service Error"}
	}

	resp.Code = userCode
	return resp, nil
}

func UpdateParty(code string, req models.Party) models.Response {
	slog.Info("updating party", "admin_code", code)
	
	db, err := database.GetDB()
	if err != nil {
		slog.Error("database connection failed", "error", err)
		return models.Response{Code: 500, Message: "Service Error"}
	}

	query := `UPDATE party SET title=?, description=?, date=?, time=?, address=? WHERE admin_code=?`
	_, err = db.Exec(query, req.Title, req.Description, req.Date, req.Time, req.Address, code)
	if err != nil {
		slog.Error("update query failed", "admin_code", code, "error", err)
		return models.Response{Code: 500, Message: "Service Error"}
	}

	return models.Response{Code: 200, Message: "success"}
}

func DeleteParty(code string) models.Response {
	slog.Info("deleting party", "admin_code", code)
	
	db, err := database.GetDB()
	if err != nil {
		slog.Error("database connection failed", "error", err)
		return models.Response{Code: 500, Message: "Service Error"}
	}

	tx, err := db.Begin()
	if err != nil {
		slog.Error("failed to begin delete transaction", "error", err)
		return models.Response{Code: 500, Message: "service error"}
	}
	defer tx.Rollback()

	var party_id int
	err = tx.QueryRow("SELECT id FROM party WHERE admin_code=?", code).Scan(&party_id)
	if err != nil {
		slog.Error("failed to find party for deletion", "admin_code", code, "error", err)
		return models.Response{Code: 500, Message: "service error"}
	}

	tables := []string{"reminders", "announcements", "guests", "posts", "party"}
	for _, table := range tables {
		query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", table, map[bool]string{table == "party": "id"}[table == "party"])
		if table != "party" {
			query = fmt.Sprintf("DELETE FROM %s WHERE party_id = ?", table)
		}
		if _, err := tx.Exec(query, party_id); err != nil {
			slog.Error("delete failed", "table", table, "party_id", party_id, "error", err)
			return models.Response{Code: 500, Message: "service error"}
		}
	}

	if err := tx.Commit(); err != nil {
		slog.Error("delete transaction commit failed", "party_id", party_id, "error", err)
		return models.Response{Code: 500, Message: "service error"}
	}
	return models.Response{Code: 200, Message: "success"}
}
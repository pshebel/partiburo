package operations

import (
	"fmt"
	"log"
	"strings"
	"database/sql"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/utils"
	"github.com/pshebel/partiburo/backend/database"
	"github.com/pshebel/partiburo/backend/notifications"

)

// takes a list of codes and returns titles for them. omits titles from 
func GetTitles(req models.TitlesRequest) (models.TitlesResponse, *models.Response) {
	log.Println("Get Titles")
	resp := models.TitlesResponse{}
	if len(req.Codes) == 0 {
		log.Println("no codes submitted")
		return resp, &models.Response{404, "no titles"}
	}
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return resp, &models.Response{500, "Service Error"}
	}

	placeholders := make([]string, len(req.Codes))
    for i := range req.Codes {
        placeholders[i] = "?"
    }
    query := fmt.Sprintf(
        "SELECT user_code, title FROM party WHERE user_code IN (%s);", 
        strings.Join(placeholders, ","),
    )
	args := make([]interface{}, len(req.Codes))
    for i, v := range req.Codes {
        args[i] = v
    }
	
	rows, err := db.Query(query, args...)
    if err != nil {
        return resp, &models.Response{500, "Service Error"}
    }
    defer rows.Close()
	results := make(map[string]string)
    for rows.Next() {
        var userCode, title string
        err := rows.Scan(&userCode, &title)
		if err != nil && err != sql.ErrNoRows {
            return resp, &models.Response{500, "Service Error"}
        }
		if err == sql.ErrNoRows {
			// omit this code from response
			continue
		}
        results[userCode] = title
    }

	resp.Titles = results
	return resp, nil
}

func GetParty(code string) (models.Party, *models.Response) {
	log.Println("GetParty")
	party := models.Party{}

	party_id := 0

	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return party, &models.Response{500, "Service Error"}
	}
	
	partyQuery := `SELECT id, date, time, address, title, description FROM party WHERE user_code = ? OR admin_code= ?`
	row := db.QueryRow(partyQuery, code, code)
	err = row.Scan(&party_id, &party.Date, &party.Time, &party.Address, &party.Title, &party.Description)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows{
			return party, &models.Response{404, "Could not find party"}
		}
		return party, &models.Response{500, "Service Error"}
	}

	query := `SELECT day_of, day_before, week_before, announcements from reminders where party_id=?`
	row = db.QueryRow(query, party_id)
	var dayOf, dayBefore, weekBefore, announcements bool
	err = row.Scan(&dayOf, &dayBefore, &weekBefore, &announcements)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows{
			return party, &models.Response{404, "Could not find party"}
		}
		return party, &models.Response{500, "Service Error"}
	}
	if (dayOf) {
		party.Reminders = append(party.Reminders, "day_of")
	}
	if (dayBefore) {
		party.Reminders = append(party.Reminders, "day_before")
	}	
	if (weekBefore) {
		party.Reminders = append(party.Reminders, "week_before")
	}	
	if (announcements) {
		party.Reminders = append(party.Reminders, "announcements")
	}

	query = `SELECT id, header, body, created_at FROM announcements WHERE party_id=?`
	rows, err := db.Query(query, party_id)
	if err != nil {
		log.Println(err)
		return party, &models.Response{500, "service error"}
	}
	party.Announcements = []models.Announcement{}
	defer rows.Close()
	for rows.Next() {
		var a models.Announcement
		err := rows.Scan(&a.ID, &a.Header, &a.Body, &a.CreatedAt)
		if err != nil {
			log.Println(err)
			return party, &models.Response{500, "service error"}
		}
		party.Announcements = append(party.Announcements, a)
	}

	return party, nil
}


func CreateParty(req models.PartyRequest) (models.PartyResponse, *models.Response) {
	log.Println("CreateParty")
	resp := models.PartyResponse{}

	confirmed, err := notifications.ConfirmEmail(req.AdminEmail) 
	if err != nil {
		log.Println(err)
		return resp, &models.Response{500, "Service Error"}
	}

	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return resp, &models.Response{500, "Service Error"}
	}

	userCode := utils.RandomString()
	adminCode := utils.RandomString()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return resp, &models.Response{500, "Service Error"}
	}

	defer tx.Commit()
	email_id := 0
	query := `SELECT id FROM email WHERE email=?`
	row := tx.QueryRow(query, req.AdminEmail)
	err = row.Scan(&email_id)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return resp, &models.Response{500, "Service Error"}
	}

	res, err := tx.Exec("INSERT INTO party (admin_email_id, user_code, admin_code, date, time, address, title, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", email_id, userCode, adminCode, req.Date, req.Time, req.Address, req.Title, req.Description)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return resp, &models.Response{500, "Service Error"}
	}

	party_id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return resp, &models.Response{500, "Service Error"}
	}

	if (len(req.Reminders) > 0) {
		var announcements, dayOf, dayBefore, weekBefore bool
		for _, r := range req.Reminders {
			switch r {
			case "day_of":
				dayOf = true
			case "day_before":
				dayBefore = true
			case "week_before":
				weekBefore = true
			case "announcements": 
				announcements = true
			}
		}

		query = `INSERT INTO reminders (party_id, day_of, day_before, week_before, announcements) VALUES (? ,? ,? ,? ,?)`
		_, err = tx.Exec(query, party_id, dayOf, dayBefore, weekBefore, announcements)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return resp, &models.Response{500, "Service Error"}
		}
	} 


	subject := fmt.Sprintf("Here are your Partiburo links for %s", req.Title)
	message := fmt.Sprintf("Guest Link: https://partiburo.com/%s\nAdmin Link: https://partiburo.com/admin/%s\n", userCode, adminCode)
	if (confirmed) {
		err := notifications.PublishEmail(req.AdminEmail, subject, message)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return resp, &models.Response{500, "Service Error"}
		}
	} else {
		_, err = tx.Exec("INSERT INTO queue (email_id, subject, body, sent, retry) VALUES (?, ?, ?, ?, ?)", email_id, subject, message, false, "daily")
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return resp, &models.Response{500, "Service Error"}
		}
	}
	resp.Code = userCode

	return resp, nil

}


func UpdateParty(code string, req models.Party) models.Response {
	log.Println("UpdateParty")
	
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return models.Response{500, "Service Error"}
	}

	query := `UPDATE party SET title=?, description=?, date=?, time=?, address=? WHERE admin_code=?`
	_, err = db.Exec(query, req.Title, req.Description, req.Date, req.Time, req.Address, code)
	if err != nil {
		log.Println(err)
		return models.Response{500, "Service Error"}
	}

	return models.Response{200, "success"}
}

func DeleteParty(code string) models.Response {
	log.Println("DeleteParty")
	
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return models.Response{500, "Service Error"}
	}

	tx, err := db.Begin()
    if err != nil {
		log.Println(err)
		return models.Response{500, "service error"}
    }

	defer tx.Rollback()

	party_id := 0
	row := tx.QueryRow("SELECT id FROM party WHERE admin_code=?", code)
	err = row.Scan(&party_id)
	if err != nil {
		log.Println(err)
		return models.Response{500, "service error"}
	}
	_, err = tx.Exec("DELETE FROM announcements WHERE party_id = ?", party_id)
	if err != nil {
		log.Println(err)
		return models.Response{500, "service error"}
    }
    _, err = tx.Exec("DELETE FROM guests WHERE party_id = ?", party_id)
	if err != nil {
		log.Println(err)
		return models.Response{500, "service error"}
    }
    _, err = tx.Exec("DELETE FROM party WHERE id = ?", party_id)
	if err != nil {
		log.Println(err)
		return models.Response{500, "service error"}
    }

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return models.Response{500, "service error"}
    }
	return models.Response{200, "success"}
}
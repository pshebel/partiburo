package operations

import (
	"fmt"
	"log"
	"strings"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/utils"
	"github.com/pshebel/partiburo/backend/database"
	"github.com/pshebel/partiburo/backend/notifications"

)

func GetTitles(req models.TitlesRequest) (models.TitlesResponse, error) {
	log.Println("Get Titles")
	resp := models.TitlesResponse{}
	if len(req.Codes) == 0 {
		return resp, nil
	}
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return resp, err
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
        return resp, fmt.Errorf("query failed: %w", err)
    }
    defer rows.Close()
	results := make(map[string]string)
    for rows.Next() {
        var userCode, title string
        err := rows.Scan(&userCode, &title)
		if err != nil {
            return resp, fmt.Errorf("scan failed: %w", err)
        }
        results[userCode] = title
    }

    // Check for errors encountered during iteration
    if err = rows.Err(); err != nil {
        return resp, err
    }
	resp.Titles = results
	return resp, nil
}

func GetParty(code string) (models.Party, error) {
	log.Println("GetParty")
	party := models.Party{}

	party_id := 0

	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return party, err
	}
	
	partyQuery := `SELECT id, date, time, address, title, description FROM party WHERE user_code = $1`
	row := db.QueryRow(partyQuery, code)
	err = row.Scan(&party_id, &party.Date, &party.Time, &party.Address, &party.Title, &party.Description)
	if err != nil {
		log.Println(err)
		return party, err
	}
	return party, nil
}


func CreateParty(req models.PartyRequest) (models.PartyResponse, error) {
	log.Println("CreateParty")
	resp := models.PartyResponse{}

	confirmed, err := notifications.ConfirmEmail(req.AdminEmail) 
	if err != nil {
		log.Println(err)
		return resp, err
	}

	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return resp, err
	}

	userCode := utils.RandomString()
	adminCode := utils.RandomString()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return resp, err
	}

	defer tx.Commit()
	email_id := 0
	query := `SELECT id FROM email WHERE email=?`
	row := tx.QueryRow(query, req.AdminEmail)
	err = row.Scan(&email_id)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return resp, err
	}

	res, err := tx.Exec("INSERT INTO party (admin_email_id, user_code, admin_code, date, time, address, title, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", email_id, userCode, adminCode, req.Date, req.Time, req.Address, req.Title, req.Description)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return resp, err
	}

	party_id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return resp, err
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
			return resp, err
		}
	} 


	subject := fmt.Sprintf("Here are your Partiburo links for %s", req.Title)
	message := fmt.Sprintf("Guest Link: https://partiburo.com/%s\nAdmin Link: https://partiburo.com/admin/%s\n", userCode, adminCode)
	if (confirmed) {
		err := notifications.PublishEmail(req.AdminEmail, subject, message)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return resp, err
		}
	} else {
		_, err = tx.Exec("INSERT INTO queue (email_id, subject, body, sent, retry) VALUES (?, ?, ?, ?, ?)", email_id, subject, message, false, "daily")
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return resp, err
		}
	}
	resp.Code = userCode

	return resp, nil

}


func UpdateParty(code string, req models.Party) (models.Response, error) {
	log.Println("UpdateParty")
	resp := models.Response{}

	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return resp, err
	}

	query := `UPDATE party SET title=?, description=?, date=?, time=?, address=? WHERE admin_code=?`
	_, err = db.Exec(query, req.Title, req.Description, req.Date, req.Time, req.Address)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	resp.Code = 200
	return resp, nil
}
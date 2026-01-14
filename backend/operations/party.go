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
	_, err = tx.Exec("INSERT INTO party (admin_email, user_code, admin_code, date, time, address, title, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", req.AdminEmail, userCode, adminCode, req.Date, req.Time, req.Address, req.Title, req.Description)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return resp, err
	}

	confirmed, err := notifications.ConfirmEmail(req.AdminEmail) 
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return resp, err
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
		_, err = tx.Exec("INSERT INTO queue (email, subject, body, sent, retry) VALUES (?, ?, ?, ?, ?)", req.AdminEmail, subject, message, false, "daily")
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return resp, err
		}
	}


	return resp, nil

}
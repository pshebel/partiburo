package operations

import (
	"log"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/notifications"
	"github.com/pshebel/partiburo/backend/database"
)

func CreateConfirm(req models.ConfirmRequest) (models.Response, error) {
	log.Println("CreateConfirm")
	resp := models.Response{}
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return resp, err
	}

	email_id := 0
	email := ""
	query := `SELECT email, id FROM email WHERE code=?`
	row := db.QueryRow(query, req.Code)
	err = row.Scan(&email, &email_id)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	confirmed := false
	query = `SELECT confirmed FROM whitelist WHERE email_id = ?`
	row = db.QueryRow(query, email_id)
	err = row.Scan(&confirmed)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	// email already confirmed
	if (confirmed) {
		resp.Code = 200
		resp.Message = "email confirmed"
		return resp, nil
	}


	query = `UPDATE whitelist SET confirmed=true WHERE email_id=? AND passcode=?`
	res, err := db.Exec(query, email_id, req.Passcode)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return resp, err
	}

	if (rowsAffected == 0) {
		resp.Code = 404
		resp.Message = "email passcode combination not found"
		return resp, nil
	}

	query = `SELECT subject, body FROM queue WHERE sent=false AND email_id=?`
	rows, err := db.Query(query, email_id)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var body, subject string

		err := rows.Scan(&subject, &body)
		if err != nil {
			log.Println(err)
			return resp, err
		}
		err = notifications.PublishEmail(email, subject, body)
		if err != nil {
			log.Println(err)
			return resp, err
		}
	}
	
	resp.Code = 200
	resp.Message = "email confirmed"
	return resp, nil
}
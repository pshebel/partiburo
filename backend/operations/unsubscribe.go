package operations


import (
	"log"
	"database/sql"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/database"
)


func CreateUnsubscribe(req models.Unsubscribe) (models.Response, error) {
	log.Println("CreateUnsubscribe")
	resp := models.Response{}
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return resp, nil
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return resp, err
	}

	defer tx.Commit()
	email_id := 0
	query := `SELECT id FROM email WHERE code=?`
	row := tx.QueryRow(query, req.EmailCode)
	err = row.Scan(&email_id)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return resp, err
	}

	if (req.All) {
		blacklist := `insert into blacklist (email_id) values (?)`
		_, err = tx.Exec(blacklist, email_id)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return resp, err
		}
		resp.Code = 200
		resp.Message = "Successfully unsubscribed from all communication"
		return resp, nil
	}

	query = `update guests set email_id=NULL from party where guests.email_id=? and party.user_code=?`
	_, err = tx.Exec(query, email_id, req.PartyCode)
	if err != nil && err != sql.ErrNoRows{
		tx.Rollback()
		log.Println(err)
		return resp, err
	}
	resp.Code = 200
	resp.Message = "Successfully unsubscribed from party reminders"
	return resp, nil
}
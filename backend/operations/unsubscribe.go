package operations


import (
	"log"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/database"
)


func CreateUnsubscribe(req models.Unsubscribe) (models.Response, error) {
	resp := models.Response{}
	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
		return resp, nil
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return resp, err
	}

	defer tx.Commit()

	if (req.All) {
		query := `update guests set email='' where email=?`
		_, err := tx.Exec(query, req.Email)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
			return resp, err
		}

		blacklist := `insert into blacklist (email) values (?)`
		_, err = tx.Exec(blacklist, req.Email)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
			return resp, err
		}
		resp.Code = 200
		resp.Message = "Successfully unsubscribed from all communication"
		return resp, nil
	}

	query := `update guests set email='' where email=? and party_id=?`
	_, err = tx.Exec(query, req.Email, req.PartyId)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return resp, err
	}
	resp.Code = 200
	resp.Message = "Successfully unsubscribed from party reminders"
	return resp, nil
}
package operations


import (
	"log"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/database"
)

func CreateAnnouncement(code string, req models.Announcement) (models.Response, error) {
	resp := models.Response{}
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return resp, err
	}
	party_id := 0
	query := `SELECT id FROM party WHERE admin_code=?`
	row := db.QueryRow(query, code)
	err = row.Scan(&party_id)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	query = `INSERT INTO announcements (party_id, header, body) VALUES (?, ?, ?)`
	_, err = db.Exec(query, party_id, req.Header, req.Body)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	resp.Code = 200
	resp.Message = "successfully created announcement"
	return resp, nil
}
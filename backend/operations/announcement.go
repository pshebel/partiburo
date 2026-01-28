package operations


import (
	"log"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/database"
)

func GetAnnouncements(code string) ([]models.Announcement, *models.Response){
	announcements := []models.Announcement{}

	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return announcements, &models.Response{500, "service error"}
	}

	query := `SELECT a.id, a.header, a.body, a.created_at FROM announcements as a LEFT JOIN party as p ON a.party_id=p.id WHERE p.user_code=? OR p.admin_code=?`
	rows, err := db.Query(query, code, code)
	if err != nil {
		log.Println(err)
		return announcements, &models.Response{500, "service error"}
	}

	defer rows.Close()
	for rows.Next() {
		var a models.Announcement
		err := rows.Scan(&a.ID, &a.Header, &a.Body, &a.CreatedAt)
		if err != nil {
			log.Println(err)
			return announcements, &models.Response{500, "service error"}
		}
		announcements = append(announcements, a)
	}

	return announcements, nil
}

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


func UpdateAnnouncement(code string, req models.Announcement) models.Response {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return models.Response{500, "service error"}
	}

	query := `
		UPDATE announcements 
		SET 
		header = ?, 
		body = ? 
		WHERE id = ? 
		AND party_id = (
			SELECT id 
			FROM party 
			WHERE admin_code = ?
		);
	`

	_, err = db.Exec(query, req.Header, req.Body, req.ID, code)
	if err != nil {
		log.Println(err)
		return models.Response{500, "Service Error"}
	}

	return models.Response{200, "success"}
}


func DeleteAnnouncement(code string, req models.Announcement) models.Response {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return models.Response{500, "service error"}
	}

	query := `
		DELETE FROM announcements 
		WHERE id = ? 
		AND party_id = (
			SELECT id 
			FROM party 
			WHERE admin_code = ?
		);
	`

	_, err = db.Exec(query, req.ID, code)
	if err != nil {
		log.Println(err)
		return models.Response{500, "Service Error"}
	}

	return models.Response{200, "success"}
}
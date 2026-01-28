package operations

import (
	"log"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/database"
)

func CreatePost(code string, req models.PostRequest) (models.Post, error) {
	log.Println("CreatePost")
	post := models.Post{}
	party_id := 0

	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return post, nil
	}
	query := `SELECT id FROM party WHERE user_code=?`
	row := db.QueryRow(query, code)
	err = row.Scan(&party_id)
	if err != nil {
		log.Println(err)
		return post, err
	}

	query = `INSERT INTO posts (body, guest_id, party_id) VALUES (?, ?, ?)`
	_, err = db.Exec(query, req.Body, req.ID, party_id)
    if err != nil {
		log.Println(err)
        return post, nil
    }

	return post, nil
}

func UpdatePost(code string, req models.Post) models.Response {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return models.Response{500, "service error"}
	}

	query := `
		UPDATE posts 
		SET
		body = ? 
		WHERE id = ? 
		AND party_id = (
			SELECT id 
			FROM party 
			WHERE user_code = ?
		);
	`
	_, err = db.Exec(query, req.Body, req.ID, code)
	if err != nil {
		log.Println(err)
		return models.Response{500, "Service Error"}
	}

	return models.Response{200, "success"}
}

func DeletePost(code string, req models.Post) models.Response {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return models.Response{500, "service error"}
	}

	query := `
		DELETE FROM posts 
		WHERE id = ? 
		AND party_id = (
			SELECT id 
			FROM party 
			WHERE user_code = ?
		);
	`
	_, err = db.Exec(query, req.ID, code)
	if err != nil {
		log.Println(err)
		return models.Response{500, "Service Error"}
	}

	return models.Response{200, "success"}
}
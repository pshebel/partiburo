package operations

import (
	"log"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/database"
)

func CreatePost(req models.PostRequest) (models.Post, error) {
	party_id := 0
	post := models.Post{}

	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
		return post, nil
	}
	query := `INSERT INTO posts (body, guest_id, party_id) VALUES (?, ?, ?)`
	res, err := db.Exec(query, req.Body, req.GuestID, party_id)
    if err != nil {
		log.Fatal(err)
        return post, nil
    }

    id, err := res.LastInsertId()
    if err != nil {
		log.Fatal(err)
        return post, err
    }

	query = `SELECT p.id, g.name, p.body, p.created_at 
	FROM posts as p 
	LEFT JOIN guests as g ON p.guest_id = g.id
	WHERE p.id=?`
	row := db.QueryRow(query, id)
	err = row.Scan(&post.ID, &post.Name, &post.Body, &post.CreatedAt)
	if err != nil {
		log.Fatal(err)
		return post, err
	}

	return post, nil
}
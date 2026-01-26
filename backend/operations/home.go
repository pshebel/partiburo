package operations

import (
	"log"
	"database/sql"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/database"
)

func GetHome(code string) (models.Home, error) {
	log.Println("GetHome")
	home := models.Home{}

	party_id := 0

	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return home, err
	}
	
	partyQuery := `SELECT id, date, time, address, title, description FROM party WHERE user_code = ? OR admin_code = ?`
	row := db.QueryRow(partyQuery, code, code)
	err = row.Scan(&party_id, &home.Date, &home.Time, &home.Address, &home.Title, &home.Description)
	if err != nil {
		log.Println(err)
		return home, err
	}

	announcementsQuery := `SELECT header, body, created_at FROM announcements where party_id = $1`
	rows, err := db.Query(announcementsQuery, party_id)
	if err != nil {
		log.Println(err)
		return home, err
	}
	defer rows.Close()

	home.Announcements = []models.Announcement{}
	for rows.Next() {
		var a models.Announcement
		err := rows.Scan(&a.Header, &a.Body, &a.CreatedAt)
		if err != nil {
			return home, err
		}
		home.Announcements = append(home.Announcements, a)
	}
	err = rows.Err()
	if err != nil {
		return home, err
	}

	guestsQuery := `
	SELECT id, name, status, plus, created_at FROM guests WHERE party_id = $1`

	rows, err = db.Query(guestsQuery, party_id)
	if err != nil {
		log.Println(err)
		return home, err
	}
	defer rows.Close()

	home.Guests = []models.Guest{}
	for rows.Next() {
		var g models.Guest
		var status sql.NullString
		err := rows.Scan(&g.ID, &g.Name, &status, &g.Plus, &g.CreatedAt)
		if err != nil {
			return home, err
		}

		if status.Valid{
			g.Status = status.String
		} else {
			g.Status = ""
		}

		if g.Status == "GOING" {
			home.Going += g.Plus + 1
		}

		home.Guests = append(home.Guests, g)
	}
	err = rows.Err()
	if err != nil {
		return home, err
	}


	postsQuery := `
		SELECT 
			gu.name, 
			po.body, 
			po.created_at 
		FROM guests as gu 
		LEFT JOIN posts as po ON gu.id = po.guest_id 
		WHERE po.party_id = $1`

	rows, err = db.Query(postsQuery, party_id)
	if err != nil {
		log.Println(err)
		return home, err
	}
	defer rows.Close()

	home.Posts = []models.Post{}
	for rows.Next() {
		var p models.Post
		err := rows.Scan(&p.Name, &p.Body, &p.CreatedAt)
		if err != nil {
			return home, err
		}
		home.Posts = append(home.Posts, p)
	}
	err = rows.Err()
	if err != nil {
		return home, err
	}

	return home, nil
}

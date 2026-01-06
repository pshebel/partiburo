package operations

import (
	"log"
	"database/sql"
	"strconv"
	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/utils"
	"github.com/pshebel/partiburo/backend/database"
)

func GetParty() (models.Party, error) {
	party := models.Party{}

	party_id := 0

	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
		return party, err
	}
	
	partyQuery := `SELECT title, description FROM party WHERE id = $1`
	row := db.QueryRow(partyQuery, party_id)
	err = row.Scan(&party.Title, &party.Description)
	if err != nil {
		log.Fatal(err)
		return party, err
	}

	announcementsQuery := `SELECT header, body, created_at FROM announcements where party_id = $1`
	rows, err := db.Query(announcementsQuery, party_id)
	if err != nil {
		log.Fatal(err)
		return party, err
	}
	defer rows.Close()

	party.Announcements = []models.Announcement{}
	for rows.Next() {
		var a models.Announcement
		err := rows.Scan(&a.Header, &a.Body, &a.CreatedAt)
		if err != nil {
			return party, err
		}
		party.Announcements = append(party.Announcements, a)
	}
	err = rows.Err()
	if err != nil {
		return party, err
	}

	guestsQuery := `
	SELECT name, status, created_at FROM guests WHERE party_id = $1`

	rows, err = db.Query(guestsQuery, party_id)
	if err != nil {
		log.Fatal(err)
		return party, err
	}
	defer rows.Close()

	party.Guests = []models.Guest{}
	for rows.Next() {
		var g models.Guest
		var status sql.NullString
		err := rows.Scan(&g.Name, &status, &g.CreatedAt)
		if err != nil {
			return party, err
		}

		if status.Valid{
			g.Status = status.String
		} else {
			g.Status = ""
		}
		party.Guests = append(party.Guests, g)
	}
	err = rows.Err()
	if err != nil {
		return party, err
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
		log.Fatal(err)
		return party, err
	}
	defer rows.Close()

	party.Posts = []models.Post{}
	for rows.Next() {
		var p models.Post
		err := rows.Scan(&p.Name, &p.Body, &p.CreatedAt)
		if err != nil {
			return party, err
		}
		party.Posts = append(party.Posts, p)
	}
	err = rows.Err()
	if err != nil {
		return party, err
	}

	return party, nil
}


func CreateParty(req models.PartyRequest) (models.PartyResponse, error) {
	resp := models.PartyResponse{}

	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
		return resp, err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return resp, err
	}

	defer tx.Commit()
	res, err := tx.Exec("INSERT INTO party (title, description) VALUES (?, ?)", req.Title, req.Description)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return resp, err
	}

	party_id, err := res.LastInsertId()
    if err != nil {
		tx.Rollback()
		log.Fatal(err)
        return resp, err
    }
	token := models.Token{
		UserID: 	"0",
		PartyId: 	strconv.FormatInt(party_id, 10),
		Role: 		"Admin",
	}

	hash, err := utils.ToHashString(token)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
        return resp, err
    }

	// _, err = tx.Exec("INSERT INTO party_admin_hash (party_id, admin_id, hash) VALUES (?, ?)", party_id, hash)
	// if err != nil {
	// 	tx.Rollback()
	// 	log.Fatal(err)
	// 	return resp, err
	// }
	
	

	resp.TokenHash = hash
	return resp, nil

}
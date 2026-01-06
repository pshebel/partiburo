package operations


import (
	"log"
	"strconv"
	"database/sql"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/database"
)


func GetGuests() ([]models.Guest, error) {
	guests := []models.Guest{}
	party_id := 0

	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
		return guests, nil
	}

	guestsQuery := `SELECT id, name, status, created_at FROM guests WHERE party_id = $1`
	
	rows, err := db.Query(guestsQuery, party_id)
	if err != nil {
		log.Fatal(err)
		return guests, err
	}
	defer rows.Close()

	for rows.Next() {
		var g models.Guest
		var status sql.NullString
		err := rows.Scan(&g.ID, &g.Name, &status, &g.CreatedAt)
		if err != nil {
			return guests, err
		}
		if status.Valid{
			g.Status = status.String
		} else {
			g.Status = ""
		}
		guests = append(guests, g)
	}
	err = rows.Err()
	if err != nil {
		return guests, err
	}
	return guests, nil
}


func CreateGuest(guest models.GuestRequest) (models.GuestResponse, error) {
	resp := models.GuestResponse{}
	party_id := 0

	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
		return resp, nil
	}

	guestQuery := `INSERT INTO guests (name, status, party_id) VALUES (?, ?, ?)`
	res, err := db.Exec(guestQuery, guest.Name, guest.Status, party_id)
    if err != nil {
		log.Fatal(err)
        return resp, nil
    }

    id, err := res.LastInsertId()
    if err != nil {
		log.Fatal(err)
        return resp, err
    }

	resp.ID = strconv.FormatInt(id, 10)
	return resp, nil
}
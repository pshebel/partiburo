package operations


import (
	"log"
	"strconv"
	"database/sql"
	"strings"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/utils"
	"github.com/pshebel/partiburo/backend/database"
	"github.com/pshebel/partiburo/backend/notifications"
)


func GetGuests() ([]models.Guest, error) {
	log.Println("GetGuests")
	guests := []models.Guest{}
	party_id := 0

	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return guests, nil
	}

	guestsQuery := `SELECT id, name, status, plus, created_at FROM guests WHERE party_id = $1`
	
	rows, err := db.Query(guestsQuery, party_id)
	if err != nil {
		log.Println(err)
		return guests, err
	}
	defer rows.Close()

	for rows.Next() {
		var g models.Guest
		var status sql.NullString
		err := rows.Scan(&g.ID, &g.Name, &status, &g.Plus, &g.CreatedAt)
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
	log.Println("CreateGuest")
	resp := models.GuestResponse{}
	party_id := 0

	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return resp, nil
	}

	guest.Email = strings.TrimSpace(guest.Email)

	if guest.Email != "" && utils.IsValidEmail(guest.Email) {
		_, err := notifications.ConfirmEmail(guest.Email)
		if err != nil {
			log.Println(err)
			return resp, nil
		}
	}


	guestQuery := `INSERT INTO guests (name, email, status, plus, party_id) VALUES (?, ?, ?, ?, ?)`
	res, err := db.Exec(guestQuery, guest.Name, guest.Email, guest.Status, guest.Plus, party_id)
    if err != nil {
		log.Println(err)
        return resp, nil
    }

    id, err := res.LastInsertId()
    if err != nil {
		log.Println(err)
        return resp, err
    }

	resp.ID = strconv.FormatInt(id, 10)
	return resp, nil
}

func UpdateGuest(guest models.UpdateGuestRequest) (models.Guest, error) {
	log.Println("UpdateGuest")
	resp := models.Guest{}
	party_id := 0

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

	if guest.Status == "GOING" || guest.Status == "NOT_GOING" || guest.Status == "MAYBE" {
		query := `UPDATE guests SET status=? WHERE party_id=? AND id=?`
		_, err := db.Exec(query, guest.Status, party_id, guest.ID)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return resp, nil
		}
	}

	guest.Email = strings.TrimSpace(guest.Email)
	if guest.Email != "" && utils.IsValidEmail(guest.Email) {
		_, err := notifications.ConfirmEmail(guest.Email)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return resp, nil
		}

		query := `UPDATE guests SET email=? WHERE party_id=? AND id=?`
		_, err = db.Exec(query, guest.Email, party_id, guest.ID)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return resp, nil
		}
	}

	query := `UPDATE guests SET plus=? WHERE party_id=? AND id=?`
	_, err = db.Exec(query, guest.Plus, party_id, guest.ID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return resp, nil
	}
	
	return resp, nil
}
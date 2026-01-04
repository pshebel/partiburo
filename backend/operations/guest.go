package operations


import (
	"fmt"
	"log"
	"strconv"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/database"
	"github.com/pshebel/partiburo/backend/utils"
)


func GetGuests() ([]models.Guest, error) {
	party_id := 0
	guests := []models.Guest{}
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
		err := rows.Scan(&g.ID, &g.Name, &g.Status, &g.CreatedAt)
		if err != nil {
			return guests, err
		}
		guests = append(guests, g)
	}
	err = rows.Err()
	if err != nil {
		return guests, err
	}
	return guests, nil
}


func CreateGuest(guest models.GuestRequest, token_hash string) (models.GuestResponse, error) {
	resp := models.GuestResponse{}

	token, err := utils.FromHashString(token_hash)
	if err != nil {
		log.Fatal(err)
		return resp, err
	}

	fmt.Printf("%v\n", token)

	if token.Role != "Admin" {
		log.Fatal("unauth")
		return resp, err
	}

	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
		return resp, nil
	}

	guestQuery := `INSERT INTO guests (name, party_id) VALUES (?, ?)`
	res, err := db.Exec(guestQuery, guest.Name, token.PartyId)
    if err != nil {
		log.Fatal(err)
        return resp, nil
    }

    id, err := res.LastInsertId()
    if err != nil {
		log.Fatal(err)
        return resp, err
    }

	guest_token := models.Token{
		UserID: strconv.FormatInt(id, 10),
		PartyId: token.PartyId,
		Role: "Guest",
	}

	hash, err := utils.ToHashString(guest_token)
	if err != nil {
		log.Fatal(err)
        return resp, err
    }

	resp.TokenHash = hash
	return resp, nil
}
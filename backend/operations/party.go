package operations

import (
	"log"

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
	
	partyQuery := `SELECT date, time, address, title, description FROM party WHERE id = $1`
	row := db.QueryRow(partyQuery, party_id)
	err = row.Scan(&party.Date, &party.Time, &party.Address, &party.Title, &party.Description)
	if err != nil {
		log.Fatal(err)
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
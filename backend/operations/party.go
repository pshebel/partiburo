package operations

import (
	"log"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/database"
)

func GetParty() (models.Party, error) {
	log.Println("GetParty")
	party := models.Party{}

	party_id := 0

	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return party, err
	}
	
	partyQuery := `SELECT date, time, address, title, description FROM party WHERE id = $1`
	row := db.QueryRow(partyQuery, party_id)
	err = row.Scan(&party.Date, &party.Time, &party.Address, &party.Title, &party.Description)
	if err != nil {
		log.Println(err)
		return party, err
	}
	return party, nil
}


func CreateParty(req models.PartyRequest) (models.PartyResponse, error) {
	log.Println("CreateParty")
	resp := models.PartyResponse{}

	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return resp, err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return resp, err
	}

	defer tx.Commit()
	_, err = tx.Exec("INSERT INTO party (date, time, address, title, description) VALUES (?, ?, ?, ?, ?)", req.Date, req.Time, req.Address, req.Title, req.Description)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return resp, err
	}

	return resp, nil

}
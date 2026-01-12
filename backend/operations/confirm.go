package operations

import (
	"log"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/database"
)

func CreateConfirm(req models.ConfirmRequest) (models.Response, error) {
	resp := models.Response{}
	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
		return resp, nil
	}

	query := `UPDATE whitelist SET confirmed=true WHERE email=? AND passcode=?`
	_, err = db.Exec(query, req.Email, req.Passcode)
	if err != nil {
		log.Fatal(err)
		return resp, nil
	}

	resp.Code = 200
	resp.Message = "email confirmed"
	return resp, nil
}
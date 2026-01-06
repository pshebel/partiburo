package transport

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/operations"
)

func GetPartyHandler(w http.ResponseWriter, r *http.Request) {
	party, err := operations.GetParty()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "party not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(party)
}

func CreatePartyHandler(w http.ResponseWriter, r *http.Request) {
	var req models.PartyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := models.Response{
			Code: 500,
			Message: "invalid request body",
		}
		json.NewEncoder(w).Encode(resp)
        return
    }

	party, err := operations.CreateParty(req)
	if err != nil {
		http.Error(w, "party not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(party)
}

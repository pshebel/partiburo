package transport

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/operations"
)

func GetPartyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    tokenHash := vars["token_hash"]
    if tokenHash == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(models.Response{
            Code:    400,
            Message: "missing token_hash",
        })
        return
    }

	party, err := operations.GetParty(tokenHash)
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

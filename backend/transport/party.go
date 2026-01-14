package transport

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/operations"
)

func PostTitlesHandler(w http.ResponseWriter, r *http.Request) {
	var req models.TitlesRequest
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		resp := models.Response{
			Code: 500,
			Message: "invalid request body",
		}
		json.NewEncoder(w).Encode(resp)
        return
    }

	titles, err := operations.GetTitles(req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "titles not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(titles)
}

func GetPartyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    code := vars["code"]
    if code == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(models.Response{
            Code:    400,
            Message: "missing code",
        })
        return
    }
	party, err := operations.GetParty(code)
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


func UpdatePartyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    code := vars["code"]
    if code == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(models.Response{
            Code:    400,
            Message: "missing code",
        })
        return
    }
	var req models.Party
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := models.Response{
			Code: 500,
			Message: "invalid request body",
		}
		json.NewEncoder(w).Encode(resp)
        return
    }

	party, err := operations.UpdateParty(code, req)
	if err != nil {
		http.Error(w, "party not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(party)
}

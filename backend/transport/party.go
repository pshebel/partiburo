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
		fmt.Println(http.StatusBadRequest)
		resp := models.Response{
			Code: http.StatusBadRequest,
			Message: "invalid request body",
		}
		json.NewEncoder(w).Encode(resp)
        return
    }

	titles, resp := operations.GetTitles(req)
	if resp != nil {
		fmt.Println(resp)
		w.WriteHeader(resp.Code)
        json.NewEncoder(w).Encode(resp)
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
            Code:    http.StatusBadRequest,
            Message: "missing code",
        })
        return
    }
	party, resp := operations.GetParty(code)
	if resp != nil {
		fmt.Println(resp)
		w.WriteHeader(resp.Code)
        json.NewEncoder(w).Encode(resp)
		return
	}

	json.NewEncoder(w).Encode(party)
}

func CreatePartyHandler(w http.ResponseWriter, r *http.Request) {
	var req models.PartyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(models.Response{
            Code:    http.StatusBadRequest,
            Message: "bad request",
        })
        return
    }

	party, resp := operations.CreateParty(req)
	if resp != nil {
		fmt.Println(resp)
		w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(resp)
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
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(models.Response{
            Code:    http.StatusBadRequest,
            Message: "invalid request body",
        })
        return
    }

	resp := operations.UpdateParty(code, req)

	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}

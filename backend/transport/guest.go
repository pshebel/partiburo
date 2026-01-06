package transport

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/operations"
)

func GetGuestsHandler(w http.ResponseWriter, r *http.Request) {
	guests, err := operations.GetGuests()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "guests not found", http.StatusNotFound)
		return
	}
	fmt.Println(guests)
	json.NewEncoder(w).Encode(guests)
}



func CreateGuestHandler(w http.ResponseWriter, r *http.Request) {
	var req models.GuestRequest
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
	resp, err := operations.CreateGuest(req)
	if err != nil {
		resp := models.Response{
			Code: 500,
			Message: "failed to create guest",
		}
		json.NewEncoder(w).Encode(resp)
        return
	}
	json.NewEncoder(w).Encode(resp)
}

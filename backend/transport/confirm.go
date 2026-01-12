package transport

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/operations"
)


func CreateConfirmHandler(w http.ResponseWriter, r *http.Request) {
	var req models.ConfirmRequest

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
	resp, err := operations.CreateConfirm(req)
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

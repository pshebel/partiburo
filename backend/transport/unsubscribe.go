package transport

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/operations"
)

func CreateUnsubscribeHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Unsubscribe
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
	resp, err := operations.CreateUnsubscribe(req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(models.Response{
            Code:    http.StatusNotFound,
            Message: "party not found",
        })
        return
	}
	json.NewEncoder(w).Encode(resp)
}

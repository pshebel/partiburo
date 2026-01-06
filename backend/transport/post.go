package transport

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/operations"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var req models.PostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := models.Response{
			Code: 500,
			Message: "invalid request body",
		}
		json.NewEncoder(w).Encode(resp)
        return
    }

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


	post, err := operations.CreatePost(req, tokenHash)
	if err != nil {
		resp := models.Response{
			Code: 500,
			Message: "failed to create guest",
		}
		json.NewEncoder(w).Encode(resp)
        return
	}
	
	json.NewEncoder(w).Encode(post)
}

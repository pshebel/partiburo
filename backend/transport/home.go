package transport

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pshebel/partiburo/backend/operations"
	"github.com/pshebel/partiburo/backend/models"

)

func GetHomeHandler(w http.ResponseWriter, r *http.Request) {
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

	home, resp := operations.GetHome(code)
	if resp != nil {
		fmt.Println(resp)
		w.WriteHeader(resp.Code)
        json.NewEncoder(w).Encode(*resp)
		return
	}

	json.NewEncoder(w).Encode(home)
}



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

	home, err := operations.GetHome(code)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "home not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(home)
}



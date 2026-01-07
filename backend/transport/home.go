package transport

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/pshebel/partiburo/backend/operations"
)

func GetHomeHandler(w http.ResponseWriter, r *http.Request) {
	home, err := operations.GetHome()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "home not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(home)
}



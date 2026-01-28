package transport


import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pshebel/partiburo/backend/models"
	"github.com/pshebel/partiburo/backend/operations"
)

func GetAnnouncementsHandler(w http.ResponseWriter, r *http.Request) {
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

	announcements, resp := operations.GetAnnouncements(code)
	if resp != nil {
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
	}
	json.NewEncoder(w).Encode(announcements)
}

func CreateAnnouncementHandler(w http.ResponseWriter, r *http.Request) {
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
	var req models.Announcement
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := models.Response{
			Code: 500,
			Message: "invalid request body",
		}
		json.NewEncoder(w).Encode(resp)
        return
    }

	ann, err := operations.CreateAnnouncement(code, req)
	if err != nil {
		resp := models.Response{
			Code: 500,
			Message: "failed to create announcement",
		}
		json.NewEncoder(w).Encode(resp)
        return
	}
	
	json.NewEncoder(w).Encode(ann)
}


func UpdateAnnouncementHandler(w http.ResponseWriter, r *http.Request) {
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
	var req models.Announcement
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := models.Response{
			Code: 500,
			Message: "invalid request body",
		}
		json.NewEncoder(w).Encode(resp)
        return
    }

	resp := operations.UpdateAnnouncement(code, req)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}

func DeleteAnnouncementHandler(w http.ResponseWriter, r *http.Request) {
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
	var req models.Announcement
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := models.Response{
			Code: 500,
			Message: "invalid request body",
		}
		json.NewEncoder(w).Encode(resp)
        return
    }

	resp := operations.DeleteAnnouncement(code, req)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}
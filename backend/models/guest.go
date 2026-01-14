package models

import (
	"time"
)


type Guest struct {
	ID			string 		`json:"id"`
	Name		string	 	`json:"name"`
	Status		string 		`json:"status"`
	Email		string 		`json:"email"`
	Plus		int			`json:"plus"`
	CreatedAt 	time.Time 	`json:"createdAt"`
}

type UpdateGuestRequest struct {
	ID 		string 	`json:"id"`
	Status 	string 	`json:"status"`
	Email 	string 	`json:"email"`
	Plus	int 	`json:"plus"`
}

type GuestRequest struct {
	Name	string 	`json:"name"`
	Email	string 	`json:"email"`
	Status 	string 	`json:"status"`
	Plus	int 	`json:"plus"`
}

type GuestResponse struct {
	ID	string `json:"id"`
}

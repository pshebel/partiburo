package models

import (
	"time"
)


type Guest struct {
	ID			string 		`json:"id"`
	Name		string	 	`json:"name"`
	Status		string 		`json:"status"`
	Phone		string 		`json:"phone"`
	CreatedAt 	time.Time 	`json:"createdAt"`
}

type UpdateGuestRequest struct {
	ID 		string `json:"id"`
	Status 	string `json:"status"`
	Phone 	string `json:"phone"`
}

type GuestRequest struct {
	Name	string `json:"name"`
	Phone	string `json:"phone"`
	Status string `json:"status"`
}

type GuestResponse struct {
	ID	string `json:"id"`
}

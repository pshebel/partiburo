package models

import (
	"time"
)

type Post struct {
	ID			string		`json:"id"`
	GuestID		string		`json:"guest_id"`
	Name		string		`json:"name"`
	Body		string		`json:"body"`
	CreatedAt 	time.Time	`json:"created_at"`
}

type PostRequest struct {
	ID		string `json:"id"`
	Body	string `json:"body"`
}
package models

import (
	"time"
)

type Post struct {
	ID			string		`json:"id"`
	Name		string		`json:"name"`
	Body		string		`json:"body"`
	CreatedAt 	time.Time	`json:"createdAt"`
}

type PostRequest struct {
	ID		string `json:"id"`
	Body	string `json:"body"`
}
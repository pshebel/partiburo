package models

import (
	"time"
)

type Announcement struct {
	ID 			string	`json:"id"`
	Header		string	`json:"header"`
	Body		string	`json:"body"`
	CreatedAt time.Time	`json:"created_at"`
}
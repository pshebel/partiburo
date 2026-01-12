package models

import (
	"time"
)



type LoginRequest struct {
	Password	string
}

type LoginResponse struct {
	Token	string
}


type Home struct {
	ID				string
	Title			string
	Description		string
	Date			string
	Time			string
	Address			string
	Announcements	[]Announcement
	Guests			[]Guest
	Posts			[]Post
	CreatedAt 		time.Time
}

type Party struct {
	ID				string
	Title			string
	Description		string
	Date			string
	Time			string
	Address			string
	CreatedAt 		time.Time
}


type PartyRequest struct {
	Title		string	`json:"title"`
	Date		string	`json:"date"`
	Time		string	`json:"time"`
	Address		string	`json:"address"`
	Description	string	`json:"description"`
}


type Token struct {
	UserID		string `json:"user_id"`
	Role	string `json:"role"`
	PartyId	string `json:"party_id"`
}

type PartyResponse struct {
	TokenHash	string	`json:"token_hash"`
}

type Announcement struct {
	ID 			string
	Header		string
	Body		string
	CreatedAt time.Time
}


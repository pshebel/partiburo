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

type Party struct {
	ID				string
	Title			string
	Description		string
	Announcements	[]Announcement
	Guests			[]Guest
	Posts			[]Post
	CreatedAt time.Time
}


type PartyRequest struct {
	Title		string
	Description	string
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

type Guest struct {
	ID			string
	Name		string
	Status		string
	CreatedAt 	time.Time
}

type GuestRequest struct {
	Name	string `json:"name"`
	Status string `json:"status"`
}

type GuestResponse struct {
	ID	string `json:"id"`
}

type Post struct {
	ID			string
	Name		string
	Body		string
	CreatedAt 	time.Time
}


type PostRequest struct {
	GuestID		string
	Body		string
}
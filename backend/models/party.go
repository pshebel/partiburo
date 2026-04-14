package models

import (
	"time"
)

type TitlesRequest struct {
	Codes []string `json:"codes"`
}

type TitlesResponse struct {
	Titles map[string]string `json:"titles"`
}

type LoginRequest struct {
	Password	string
}

type LoginResponse struct {
	Token	string
}

type Home struct {
	Party 			Party 			`json:"party"`
	Announcements	[]Announcement 	`json:"announcements"`
	Going			int				`json:"going"`
	Guests			[]Guest			`json:"guests"`
	Posts			[]Post			`json:"posts"`
}

type Party struct {
	ID				string		`json:"id"`
	AdminEmail		string		`json:"admin_email"`
	Title			string		`json:"title"`
	Description		string		`json:"description"`
	Date			string		`json:"date"`
	Time			string		`json:"time"`
	Address			string		`json:"address"`
	Reminders		[]string	`json:"reminders"`
	Announcements	[]Announcement `json:"announcements"`
	CreatedAt 		time.Time	`json:"created_at"`
}


type PartyRequest struct {
	Title		string	`json:"title"`
	AdminEmail	string 	`json:"admin_email"`
	Date		string	`json:"date"`
	Time		string	`json:"time"`
	Address		string	`json:"address"`
	Description	string	`json:"description"`
	Reminders 	[]string `json:"reminders"`
}


type Token struct {
	UserID		string `json:"user_id"`
	Role	string `json:"role"`
	PartyId	string `json:"party_id"`
}

type PartyResponse struct {
	Code	string	`json:"code"`
}




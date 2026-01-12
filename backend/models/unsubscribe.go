package models

type Unsubscribe struct {
	PartyId	int		`json:"party_id"`
	Email	string	`json:"email"`
	All		bool	`json:"all"`
}
package models

type Unsubscribe struct {
	PartyCode	string	`json:"party_code"`
	EmailCode	string	`json:"email_code"`
	All			bool	`json:"all"`
}
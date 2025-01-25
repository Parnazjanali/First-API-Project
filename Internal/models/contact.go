package models

type Contact struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	EmailAddress string `json:"email_address"`
}

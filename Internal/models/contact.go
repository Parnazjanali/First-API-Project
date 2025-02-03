package models

type Contact struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name" gorm:"not null"`
	Phone        string `json:"phone" gorm:"unique;not null"`
	EmailAddress string `json:"email_address" gorm:"unique;not null"`
}

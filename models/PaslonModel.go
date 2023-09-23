package models

import "time"

type Paslon struct {
	ID        int       `json:"id" gorm:"primaryKey:autoIncrement"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	Visi      string    `json:"visi" gorm:"type: varchar(255)"`
	Parties   []Party   `json:"parties" gorm:"many2many:paslon_parties"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PaslonResponse struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Visi    string  `json:"visi"`
	Parties []Party `json:"parties"`
}

func (PaslonResponse) TableName() string {
	return "Paslon"
}

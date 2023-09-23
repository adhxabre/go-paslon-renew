package models

import "time"

type Party struct {
	ID        int       `json:"id" gorm:"primaryKey:autoIncrement"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	PaslonID  int       `json:"paslonId"` // Tambahkan field ini
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PartyResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (Party) TableName() string {
	return "parties"
}

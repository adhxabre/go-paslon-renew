package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type Party struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey:autoIncrement"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	PaslonID  uuid.UUID `json:"paslon_id"` // Tambahkan field ini
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PartyResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (Party) TableName() string {
	return "parties"
}

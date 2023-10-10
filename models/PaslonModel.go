package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type Paslon struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey:autoIncrement"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	Vision    string    `json:"vision" gorm:"type: varchar(255)"`
	Parties   []Party   `json:"parties" gorm:"many2many:paslon_parties"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PaslonResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Vision  string    `json:"vision"`
	Parties []Party   `json:"parties"`
}

func (PaslonResponse) TableName() string {
	return "Paslon"
}

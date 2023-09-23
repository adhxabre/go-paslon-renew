package models

import "time"

type Vote struct {
	ID        int       `json:"id" gorm:"primaryKey:autoIncrement"`
	PaslonID  int       `json:"paslon_id"`
	Paslon    Paslon    `gorm:"foreignKey:PaslonID"`
	VoterName string    `json:"voter_name" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type Vote struct {
	ID        uuid.UUID        `json:"id" gorm:"primaryKey:autoIncrement"`
	PaslonID  uuid.UUID        `json:"paslon_id"`
	Paslon    Paslon           `gorm:"foreignKey:PaslonID"`
	UserID    uuid.UUID        `json:"user_id"`
	User      UserVoteResponse `json:"user"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type VoteResponse struct {
	ID       uuid.UUID `json:"id"`
	PaslonID uuid.UUID `json:"paslon_id"`
	Paslon   Paslon    `gorm:"foreignKey:PaslonID"`
}

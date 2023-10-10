package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type Profile struct {
	ID        uuid.UUID           `json:"id" gorm:"primaryKey:autoIncrement"`
	Phone     string              `json:"phone" gorm:"type: varchar(255)"`
	Gender    string              `json:"gender" gorm:"type: varchar(255)"`
	Address   string              `json:"address" gorm:"type: varchar(255)"`
	UserID    uuid.UUID           `json:"user_id"`
	User      UserProfileResponse `json:"user"`
	CreatedAt time.Time           `json:"-"`
	UpdatedAt time.Time           `json:"-"`
}

type ProfileResponse struct {
	Phone   string    `json:"phone"`
	Gender  string    `json:"gender"`
	Address string    `json:"address"`
	UserID  uuid.UUID `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profile"
}

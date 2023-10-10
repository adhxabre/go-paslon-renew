package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID       `json:"id"`
	Name      string          `json:"name" gorm:"type: varchar(255)"`
	Email     string          `json:"email" gorm:"type: varchar(255)"`
	Password  string          `json:"-" gorm:"type: varchar(255)"`
	Role      string          `json:"role" gorm:"type: varchar(255)"`
	IsVoted   bool            `json:"is_voted"`
	Vote      VoteResponse    `json:"vote"`
	Profile   ProfileResponse `json:"profile"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
}

type UserProfileResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UserVoteResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (UserProfileResponse) TableName() string {
	return "users"
}

func (UserVoteResponse) TableName() string {
	return "users"
}

package dtos

import (
	"time"

	"github.com/google/uuid"
)

type PartyRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type PartyResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PartyResponseToPaslon struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

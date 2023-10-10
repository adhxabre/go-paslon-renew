package dtos

import (
	"time"

	"github.com/google/uuid"
)

type PaslonRequest struct {
	Name   string      `json:"name" form:"name" validate:"required"`
	Vision string      `json:"visi" form:"visi" validate:"required"`
	Party  []uuid.UUID `json:"party" form:"party"`
}

type PaslonUpdate struct {
	Name   string      `json:"name" form:"name"`
	Vision string      `json:"visi" form:"visi"`
	Party  []uuid.UUID `json:"party" form:"party"`
}

type PaslonResponse struct {
	ID        uuid.UUID               `json:"id"`
	Name      string                  `json:"name"`
	Vision    string                  `json:"visi"`
	Parties   []PartyResponseToPaslon `json:"party"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
}

type PaslonResponseToVote struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Vision string    `json:"visi"`
}

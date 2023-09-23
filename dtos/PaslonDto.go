package dtos

import (
	"time"
)

type PaslonRequest struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Visi  string `json:"visi" form:"visi" validate:"required"`
	Party []int  `json:"party" form:"party"`
}

type PaslonUpdate struct {
	Name  string `json:"name" form:"name"`
	Visi  string `json:"visi" form:"visi"`
	Party []int  `json:"party" form:"party"`
}

type PaslonResponse struct {
	ID        int                     `json:"id"`
	Name      string                  `json:"name"`
	Visi      string                  `json:"visi"`
	Parties   []PartyResponseToPaslon `json:"party"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
}

type PaslonResponseToVote struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Visi string `json:"visi"`
}

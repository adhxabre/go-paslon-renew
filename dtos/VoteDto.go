package dtos

import uuid "github.com/google/uuid"

type VoteRequest struct {
	PaslonID uuid.UUID `json:"paslon_id" form:"paslon_id" validate:"required"`
	UserID   uuid.UUID `json:"user_id" form:"user_id" validate:"required"`
}

type VoteResponse struct {
	ID     uuid.UUID            `json:"id"`
	Paslon PaslonResponseToVote `json:"paslon"`
	User   UserResponseToVote   `json:"user"`
}

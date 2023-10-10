package dtos

import (
	"backend-micro-feature/models"

	"github.com/google/uuid"
)

type ProfileRequest struct {
	Phone   string    `json:"phone" form:"phone" validate:"required"`
	Gender  string    `json:"gender" form:"gender" validate:"required"`
	Address string    `json:"address" form:"address" validate:"required"`
	UserID  uuid.UUID `json:"user_id"`
}

type ProfileResponse struct {
	Phone   string                     `json:"phone" form:"phone" validate:"required"`
	Gender  string                     `json:"gender" form:"gender" validate:"required"`
	Address string                     `json:"address" form:"address" validate:"required"`
	UserID  uuid.UUID                  `json:"user_id"`
	User    models.UserProfileResponse `json:"user"`
}

package services

import (
	"backend-micro-feature/models"

	"gorm.io/gorm"
)

type PartyServices interface {
	FindParties() ([]models.Party, error)
	GetParty(ID int) (models.Party, error)
	CreateParty(Party models.Party) (models.Party, error)
	UpdateParty(Party models.Party) (models.Party, error)
	DeleteParty(Party models.Party) (models.Party, error)
}

func ServiceParty(db *gorm.DB) *service {
	return &service{db}
}

func (s *service) FindParties() ([]models.Party, error) {
	var Parties []models.Party
	err := s.db.Find(&Parties).Error

	return Parties, err
}

func (s *service) GetParty(ID int) (models.Party, error) {
	var Party models.Party
	err := s.db.First(&Party, ID).Error

	return Party, err
}

func (s *service) CreateParty(Party models.Party) (models.Party, error) {
	err := s.db.Create(&Party).Error

	return Party, err
}

func (s *service) UpdateParty(Party models.Party) (models.Party, error) {
	err := s.db.Save(&Party).Error

	return Party, err
}

func (s *service) DeleteParty(Party models.Party) (models.Party, error) {
	err := s.db.Delete(&Party).Error

	return Party, err
}

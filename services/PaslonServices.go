package services

import (
	"backend-micro-feature/models"

	"gorm.io/gorm"
)

type PaslonServices interface {
	FindPaslons() ([]models.Paslon, error)
	GetPaslon(ID int) (models.Paslon, error)
	CreatePaslon(paslon models.Paslon) (models.Paslon, error)
	UpdatePaslon(paslon models.Paslon) (models.Paslon, error)
	DeletePaslon(paslon models.Paslon) (models.Paslon, error)
}

func ServicePaslon(db *gorm.DB) *service {
	return &service{db}
}

func (s *service) FindPaslons() ([]models.Paslon, error) {
	var paslons []models.Paslon
	err := s.db.Preload("Parties").Find(&paslons).Error

	return paslons, err
}

func (s *service) GetPaslon(ID int) (models.Paslon, error) {
	var paslon models.Paslon
	err := s.db.Preload("Parties").First(&paslon, ID).Error

	return paslon, err
}

func (s *service) CreatePaslon(paslon models.Paslon) (models.Paslon, error) {
	if err := s.db.Create(&paslon).Error; err != nil {
		return models.Paslon{}, err
	}

	if err := s.db.Model(&paslon).Preload("Parties").Find(&paslon).Error; err != nil {
		return models.Paslon{}, err
	}

	return paslon, nil
}

func (s *service) UpdatePaslon(paslon models.Paslon) (models.Paslon, error) {
	err := s.db.Save(&paslon).Error

	return paslon, err
}

func (s *service) DeletePaslon(paslon models.Paslon) (models.Paslon, error) {
	err := s.db.Delete(&paslon).Error

	return paslon, err
}

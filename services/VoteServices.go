package services

import (
	"backend-micro-feature/models"

	"gorm.io/gorm"
)

type VoteService interface {
	FindVotes() ([]models.Vote, error)
	GetVoteByPaslonID(ID int) ([]models.Vote, error)
	CreateVote(Vote models.Vote) (models.Vote, error)
}

func ServiceVote(db *gorm.DB) *service {
	return &service{db}
}

func (s *service) FindVotes() ([]models.Vote, error) {
	var votes []models.Vote
	err := s.db.Preload("Paslon").Find(&votes).Error

	return votes, err
}

func (s *service) GetVoteByPaslonID(ID int) ([]models.Vote, error) {
	var vote []models.Vote
	err := s.db.Where("paslon_id = ?", ID).Preload("Paslon").Find(&vote).Error

	return vote, err
}

func (s *service) CreateVote(vote models.Vote) (models.Vote, error) {
	if err := s.db.Create(&vote).Error; err != nil {
		return models.Vote{}, err
	}

	if err := s.db.Model(&vote).Preload("Paslon").Find(&vote).Error; err != nil {
		return models.Vote{}, err
	}

	return vote, nil
}

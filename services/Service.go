package services

import "gorm.io/gorm"

type service struct {
	db *gorm.DB
}

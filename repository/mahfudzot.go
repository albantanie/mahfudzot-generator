package repository

import (
	"gorm.io/gorm"
)

type MahfudzotRepository interface {
}

type mahfudzotRepository struct {
	db *gorm.DB
}

func NewMahfudzotRepository(db *gorm.DB) *mahfudzotRepository {
	return &mahfudzotRepository{db: db}
}

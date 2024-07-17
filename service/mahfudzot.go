package service

import (
	"github.com/albantanie/mahfudzot-generator/repository"
)

type MahfudzotService struct {
	mahfudzotRepo repository.MahfudzotRepository
}

func NewMahfudzotService(mahfudzotRepo repository.MahfudzotRepository) *MahfudzotService {
	return &MahfudzotService{mahfudzotRepo}
}

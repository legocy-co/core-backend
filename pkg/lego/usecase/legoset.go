package usecase

import (
	"legocy-go/pkg/lego/repository"
)

type LegoSetUseCase struct {
	legoSetRepo repository.LegoSetRepository
}

func NewLegoSetUseCase(legoSetRepo repository.LegoSetRepository) *LegoSetUseCase {
	return &LegoSetUseCase{
		legoSetRepo: legoSetRepo,
	}
}

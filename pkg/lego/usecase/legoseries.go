package usecase

import (
	"legocy-go/pkg/lego/repository"
)

type LegoSeriesUseCase struct {
	legoSeriesRepo repository.LegoSeriesRepository
}

func NewLegoSeriesUseCase(r repository.LegoSeriesRepository) *LegoSeriesUseCase {
	return &LegoSeriesUseCase{
		legoSeriesRepo: r,
	}
}

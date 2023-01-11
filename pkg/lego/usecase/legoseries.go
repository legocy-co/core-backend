package usecase

import (
	r "legocy-go/pkg/lego/repository"
)

type LegoSeriesUseCase struct {
	legoSeriesRepo r.LegoSeriesRepository
}

func NewLegoSeriesUseCase(repo r.LegoSeriesRepository) *LegoSeriesUseCase {
	return &LegoSeriesUseCase{
		legoSeriesRepo: repo,
	}
}

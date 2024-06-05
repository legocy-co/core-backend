package filters

type LegoSetFilterCriteria struct {
	NpiecesGTE   *int
	NpiecesLTE   *int
	SeriesIDs    []int
	SetNumbers   []int
	Name         *string
	ReleaseYears []int
}

func NewLegoSetFilterCriteria(
	NpiecesGTE *int,
	NpiecesLTE *int,
	SeriesIDs []int,
	SetNumbers []int,
	Name *string,
	ReleaseYears []int,
) *LegoSetFilterCriteria {
	return &LegoSetFilterCriteria{
		NpiecesGTE:   NpiecesGTE,
		NpiecesLTE:   NpiecesLTE,
		SeriesIDs:    SeriesIDs,
		SetNumbers:   SetNumbers,
		Name:         Name,
		ReleaseYears: ReleaseYears,
	}
}

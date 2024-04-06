package filters

type LegoSetFilterCriteria struct {
	NpiecesGTE *int
	NpiecesLTE *int
	SeriesIDs  []int
	SetNumbers []int
	Name       *string
}

func NewLegoSetFilterCriteria(
	NpiecesGTE *int,
	NpiecesLTE *int,
	SeriesIDs []int,
	SetNumbers []int,
	Name *string,
) *LegoSetFilterCriteria {
	return &LegoSetFilterCriteria{
		NpiecesGTE: NpiecesGTE,
		NpiecesLTE: NpiecesLTE,
		SeriesIDs:  SeriesIDs,
		SetNumbers: SetNumbers,
		Name:       Name,
	}
}

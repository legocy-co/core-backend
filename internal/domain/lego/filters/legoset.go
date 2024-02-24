package filters

type LegoSetFilterCriteria struct {
	NpiecesGTE *int
	NpiecesLTE *int
	SeriesIDs  *[]int
	SetNumbers *[]string
	Name       *string
}

func NewLegoSetFilterCriteria(
	NpiecesGTE *int,
	NpiecesLTE *int,
	SeriesIDs *[]int,
	SetNumbers *[]string,
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

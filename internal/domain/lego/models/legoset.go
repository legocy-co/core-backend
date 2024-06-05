package lego

type LegoSet struct {
	ID          int
	Number      int
	Name        string
	NPieces     int
	ReleaseYear int
	Series      LegoSeries
	Images      []*LegoSetImage
}

type LegoSetValueObject struct {
	Number      int
	Name        string
	NPieces     int
	SeriesID    int
	ReleaseYear int
}

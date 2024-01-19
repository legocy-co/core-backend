package lego

type LegoSet struct {
	ID      int
	Number  int
	Name    string
	NPieces int
	Series  LegoSeries
	Images  []*LegoSetImage
}

type LegoSetValueObject struct {
	Number   int    `json:"number"`
	Name     string `json:"name"`
	NPieces  int    `json:"n_pieces"`
	SeriesID int    `json:"series_id"`
}

package lego

type LegoSet struct {
	ID      int        `json:"id"`
	Number  int        `json:"number"`
	Name    string     `json:"name"`
	NPieces int        `json:"n_pieces"`
	Series  LegoSeries `json:"series"`
}

// LegoSetValueObject - Same model but without ID and
// foreign instance, pk of series stored instead
// Used when receiving request
type LegoSetValueObject struct {
	Number   int    `json:"number"`
	Name     string `json:"name"`
	NPieces  int    `json:"n_pieces"`
	SeriesID int    `json:"series_id"`
}

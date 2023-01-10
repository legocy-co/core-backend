package models

type LegoSet struct {
	ID      int        `json:"id"`
	Number  int        `json:"number"`
	Name    string     `json:"name"`
	NPieces int        `json:"n_pieces"`
	Series  LegoSeries `json:"series"`
}

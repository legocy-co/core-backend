package lego

import models "legocy-go/pkg/lego/models"

type LegoSetRequest struct {
	Name         string `json:"name"`
	Number       int    `json:"number"`
	NPieces      int    `json:"n_pieces"`
	LegoSeriesID int    `json:"series_id"`
}

func (l *LegoSetRequest) ToLegoSeriesBasic() *models.LegoSetBasic {
	return &models.LegoSetBasic{
		Name:     l.Name,
		Number:   l.Number,
		NPieces:  l.NPieces,
		SeriesID: l.LegoSeriesID,
	}
}

type LegoSetResponse struct {
	ID      int                `json:"id"`
	Name    string             `json:"name"`
	Number  int                `json:"number"`
	NPieces int                `json:"n_pieces"`
	Series  LegoSeriesResponse `json:"series"`
}

func GetLegoSetResponse(m *models.LegoSet) LegoSetResponse {
	return LegoSetResponse{
		ID:      m.ID,
		Name:    m.Name,
		Number:  m.Number,
		NPieces: m.NPieces,
		Series:  GetLegoSeriesResponse(&m.Series),
	}
}

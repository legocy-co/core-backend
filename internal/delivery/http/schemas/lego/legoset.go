package lego

import (
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
)

type LegoSetRequest struct {
	Name         string `json:"name"`
	Number       int    `json:"number"`
	NPieces      int    `json:"n_pieces"`
	LegoSeriesID int    `json:"series_id"`
}

func (l *LegoSetRequest) ToLegoSeriesValueObject() *models.LegoSetValueObject {
	return &models.LegoSetValueObject{
		Name:     l.Name,
		Number:   l.Number,
		NPieces:  l.NPieces,
		SeriesID: l.LegoSeriesID,
	}
}

type LegoSetResponse struct {
	ID      int                    `json:"id"`
	Name    string                 `json:"name"`
	Number  int                    `json:"number"`
	NPieces int                    `json:"n_pieces"`
	Series  LegoSeriesResponse     `json:"series"`
	Images  []LegoSetImageResponse `json:"images"`
}

func GetLegoSetResponse(m *models.LegoSet) LegoSetResponse {
	return LegoSetResponse{
		ID:      m.ID,
		Name:    m.Name,
		Number:  m.Number,
		NPieces: m.NPieces,
		Series:  GetLegoSeriesResponse(&m.Series),
		Images:  GetLegoSetImagesResponse(m.Images),
	}
}

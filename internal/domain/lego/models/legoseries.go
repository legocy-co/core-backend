package lego

type LegoSeries struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LegoSeriesBasic struct {
	Name string `json:"name"`
}

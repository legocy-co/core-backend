package marketplace

type MarketItemCreatedUpdated struct {
	ID          int     `json:"id"`
	LegoSetID   int     `json:"legoSetID"`
	Price       float32 `json:"price"`
	Location    string  `json:"location"`
	Status      string  `json:"status"`
	SetState    string  `json:"setState"`
	Description string  `json:"description"`
}

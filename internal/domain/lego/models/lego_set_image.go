package lego

type LegoSetImageValueObject struct {
	LegoSetID int `validate:"required"`
	ImageURL  string
	IsMain    bool `validate:"default=false"`
}

func NewLegoSetImageValueObject(legoSetID int, imageURL string, isMain bool) (*LegoSetImageValueObject, error) {
	return &LegoSetImageValueObject{
		LegoSetID: legoSetID,
		ImageURL:  imageURL,
		IsMain:    isMain,
	}, nil
}

type LegoSetImage struct {
	ID        int
	LegoSetID int `validate:"required"`
	ImageURL  string
	IsMain    bool `validate:"default=false"`
}

func NewLegoSetImage(id int, legoSetID int, imageURL string, isMain bool) (*LegoSetImage, error) {
	return &LegoSetImage{
		ID:        id,
		LegoSetID: legoSetID,
		ImageURL:  imageURL,
		IsMain:    isMain,
	}, nil
}

package marketplace

type ImageUploadResponse struct {
	OK       bool   `json:"ok"`
	ImageURL string `json:"imageURL"`
}

func NewImageUploadResponse(imageURL string) *ImageUploadResponse {
	return &ImageUploadResponse{
		OK:       imageURL != "",
		ImageURL: imageURL,
	}
}

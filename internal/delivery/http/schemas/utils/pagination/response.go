package pagination

type PageResponse[T any] struct {
	Data []T              `json:"data"`
	Meta PageMetaResponse `json:"meta"`
}

type PageMetaResponse struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func GetPageResponse[T any](data []T, total, limit, offset int) PageResponse[T] {
	return PageResponse[T]{
		Data: data,
		Meta: GetPageMetaResponse(total, limit, offset),
	}
}

func GetPageMetaResponse(total, limit, offset int) PageMetaResponse {
	return PageMetaResponse{
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}
}

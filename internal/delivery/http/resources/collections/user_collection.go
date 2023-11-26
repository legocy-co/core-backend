package collections

import (
	"legocy-go/internal/delivery/http/resources/users"
	"legocy-go/internal/domain/collections/models"
)

type UserLegoSetCollectionResponse struct {
	User users.UserDetailResponse    `json:"user"`
	Sets []CollectionLegoSetResponse `json:"collection_sets"`
}

func GetUserLegoCollectionResponse(collection models.LegoCollection) UserLegoSetCollectionResponse {

	setsResponses := make([]CollectionLegoSetResponse, 0, len(collection.Sets))
	for _, set := range collection.Sets {
		setsResponses = append(setsResponses, GetCollectionLegoSetResponse(set))
	}

	return UserLegoSetCollectionResponse{
		User: users.GetUserDetailResponse(&collection.User),
		Sets: setsResponses,
	}
}

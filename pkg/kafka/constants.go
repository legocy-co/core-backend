package kafka

import "errors"

const (
	HealthcheckTopic             = "legocy.healthcheck"
	MarketItemUpdatesTopic       = "legocy.marketItems.updates.json"
	MarketItemDeletedTopic       = "legocy.marketItems.delete.json"
	UserImagesDeletedTopic       = "legocy.users.images.delete.json"
	MarketItemImagesDeletedTopic = "legocy.marketItems.images.delete.json"
	LegoSetImagesDeletedTopic    = "legocy.legoSets.images.delete.json"
	UserCreatedTopic             = "legocy.backend.user.created.json"
	UserUpdatedTopic             = "legocy.backend.user.updated.json"
)

var ErrUnjsonableData = errors.New("cannot parse data as JSON")

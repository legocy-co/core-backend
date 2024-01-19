package kafka

import "errors"

const (
	HEALTHCHECK_TOPIC                = "legocy.healthcheck"
	USER_UPDATES_TOPIC               = "legocy.users.updates.json"
	MARKET_ITEM_UPDATES_TOPIC        = "legocy.marketItems.updates.json"
	USER_IMAGES_DELETED_TOPIC        = "legocy.users.images.delete.json"
	MARKET_ITEM_IMAGES_DELETED_TOPIC = "legocy.marketItems.images.delete.json"
	LEGO_SET_IMAGES_DELETED_TOPIC    = "legocy.legoSets.images.delete.json"
)

var ErrUnjsonableData = errors.New("cannot parse data as JSON")

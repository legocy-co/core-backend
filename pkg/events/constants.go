package events

import "errors"

const (
	HEALTHCHECK_TOPIC         = "legocy.healthcheck"
	USER_UPDATES_TOPIC        = "legocy.users.updates.json"
	MARKET_ITEM_UPDATES_TOPIC = "legocy.marketItems.updates.json"
)

var ErrUnjsonableData = errors.New("cannot parse data as JSON")

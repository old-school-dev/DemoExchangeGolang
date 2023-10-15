package api

import (
	"demo-strategy/exchange/demox/exchange"
)

type API struct {
	Key    string
	Secret string
	*exchange.Exchange
}

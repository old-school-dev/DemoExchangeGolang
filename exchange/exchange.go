package exchange

import (
	"demo-strategy/exchange/demox"
	"demo-strategy/exchange/demox/pair"
	"demo-strategy/exchange/demox/exchange"
	"demo-strategy/exchange/demox/order"
)

type IExchange interface {
	GetPairs() []pair.Pair
	GetPrice(symbol string) float64
	GetOrders() order.Orders

}

func NewExchange() *exchange.Exchange{
	return demox.NewExchange()
}

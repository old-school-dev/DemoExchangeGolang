package exchange

import (
	_"log"
	"demox/exchange/demox"
)

type IExchange interface {
	GetPairs() []demox.Pair
	GetOrderDetail(id uint64) demox.Order
	GetPrice(symbol string) float64
	CreateBuyOrder(symbol string, price float64) *demox.Order
	CreateSellOrder(symbol string, price float64) *demox.Order
	CheckMatchOrder(symbol string) []uint64
}

func NewExchange(apiKey, apiSecret string) IExchange{
	return demox.NewExchange(apiKey,apiSecret)
}

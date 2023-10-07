package exchange

import (
	"demo-strategy/exchange/demox"
	"demo-strategy/exchange/demox/models"
)

type IExchange interface {
	GetPairs() []models.Pair
	GetOrderDetail(id uint64) models.Order
	GetPrice(symbol string) float64
}

func NewExchange() *models.Exchange{
	return demox.NewExchange()
}

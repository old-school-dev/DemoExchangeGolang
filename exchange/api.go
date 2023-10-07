package exchange

import (
	"demo-strategy/exchange/demox"
	"demo-strategy/exchange/demox/models"
)

type IAPI interface {
	IExchange
	CreateBuyOrder(symbol string, price float64) *models.Order
	CreateSellOrder(symbol string, price float64) *models.Order

}

func NewAPI(key, secret string, exchange *models.Exchange) IAPI {
	return demox.NewAPI(key, secret, exchange)
}

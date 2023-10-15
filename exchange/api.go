package exchange

import (
	"demo-strategy/exchange/demox"
	"demo-strategy/exchange/demox/exchange"
	"demo-strategy/exchange/demox/order"
	"demo-strategy/exchange/demox/wallet"
)

type IAPI interface {
	IExchange
	CreateBuyOrder(pair string, price float64, amount float64) *order.Order
	CreateSellOrder(pair string, price float64, amount float64) *order.Order
	Deposit(symbol string, amount float64) float64
	GetBalances() wallet.Wallet
	GetOrderDetail(id uint64) order.Order
	GetOpenOrders() order.Orders
}

func NewAPI(key, secret string, exchange *exchange.Exchange) IAPI {
	return demox.NewAPI(key, secret, exchange)
}

package exchange

import (
	"demo-strategy/constant"
	"demo-strategy/exchange/demox/order"
	"demo-strategy/exchange/demox/pair"
	"demo-strategy/exchange/demox/wallet"
)

type Exchange struct {
	Pairs   map[string]*pair.Pair
	Orders  order.Orders
	Wallets wallet.Wallets
}

func (e *Exchange) GetPairs() []pair.Pair {
	pairs := []pair.Pair{}
	for _, pair := range e.Pairs {
		pairs = append(pairs, *pair)
	}
	return pairs
}
func (e Exchange) GetPrice(pair string) float64 {
	price := e.Pairs[pair].GetPrice()
	e.CheckMatchOrder(pair)
	return price
}

func (e *Exchange) CreateWallet(wallet wallet.Wallet) wallet.Wallet {
	e.Wallets = append(e.Wallets, wallet)
	return wallet
}

func (e *Exchange) CheckMatchOrder(pair string) []uint64 {
	ids := []uint64{}
	for index, order := range e.Orders {
		if order.Status == constant.STATUS_PENDING {
			if order.Side == constant.SIDE_BUY && order.Price >= e.Pairs[pair].Price {
				e.Orders[index].Status = constant.STATUS_MATCH
				amtAdjust := order.Amount
				currency := e.Pairs[pair].Currency
				e.AdjustWalletBalance(order.Signer, currency, amtAdjust)
				ids = append(ids, order.Id)
			} else if order.Side == constant.SIDE_SELL && order.Price <= e.Pairs[pair].Price {
				e.Orders[index].Status = constant.STATUS_MATCH
				amtAdjust := order.Amount * order.Price
				currency := e.Pairs[pair].Currency
				e.AdjustWalletBalance(order.Signer, currency, amtAdjust)
				ids = append(ids, order.Id)
			}
		}
	}
	return ids
}

func (e *Exchange) AdjustWalletBalance(walletAddress string, currency string, amount float64) {
	for i, w := range e.Wallets {
		if w.ApiKey == walletAddress {
			e.Wallets[i].Balances[currency] += amount
		}
	}
}

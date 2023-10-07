package models

type Exchange struct {
	Pairs   map[string]*Pair
	Orders  Orders
	Wallets Wallets
}

func (e *Exchange) GetPairs() []Pair {
	pairs := []Pair{}
	for _, pair := range e.Pairs {
		pairs = append(pairs, *pair)
	}
	return pairs
}

func (e *Exchange) GetOrderDetail(id uint64) Order {
	for _, order := range e.Orders {
		if order.Id == id {
			return order
		}
	}
	return Order{}
}

func (e Exchange) GetPrice(symbol string) float64 {
	price := e.Pairs[symbol].AdjustPrice()
	e.checkMatchOrder(symbol)
	return price
}

func (e *Exchange) CreateWallet(wallet Wallet) Wallet {
	e.Wallets = append(e.Wallets, wallet)
	return wallet
}

func (e *Exchange) checkMatchOrder(symbol string) []uint64 {
	ids := []uint64{}
	for index, order := range e.Orders {
		if order.Status == STATUS_PENDING {
			if order.Side == SIDE_BUY && order.Price >= e.Pairs[symbol].Price {
				e.Orders[index].Status = STATUS_MATCH
				ids = append(ids, order.Id)
			} else if order.Side == SIDE_SELL && order.Price <= e.Pairs[symbol].Price {
				e.Orders[index].Status = STATUS_MATCH
				ids = append(ids, order.Id)
			}
		}
	}
	return ids
}

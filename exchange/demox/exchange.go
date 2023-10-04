package demox

import "log"

type Exchange struct {
	Pairs     map[string]*Pair
	ApiKey    string
	ApiSecret string
	Orders    []Order
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

func (e *Exchange) getNewId() uint64 {
	if len(e.Orders) == 0 {
		return 0
	}
	return e.Orders[len(e.Orders)-1].Id + 1
}

func (e Exchange) GetPrice(symbol string) float64 {
	price := e.Pairs[symbol].AdjustPrice()
	matchIds := e.CheckMatchOrder(symbol)
	log.Println(e.Pairs[symbol].Price, matchIds)
	return price
}

func (e *Exchange) CreateBuyOrder(symbol string, price float64) *Order {
	id := e.getNewId()
	newOrder := Order{
		Id:     id,
		Symbol: symbol,
		Type:   ORDER_LIMIT,
		Side:   SIDE_BUY,
		Price:  price,
		Status: STATUS_PENDING,
	}
	e.Orders = append(e.Orders, newOrder)
	matchIds := e.CheckMatchOrder(symbol)
	log.Println(matchIds)
	return &newOrder
}

func (e *Exchange) CreateSellOrder(symbol string, price float64) *Order {
	id := e.getNewId()
	newOrder := Order{
		Id:     id,
		Symbol: symbol,
		Type:   ORDER_LIMIT,
		Side:   SIDE_SELL,
		Price:  price,
		Status: STATUS_PENDING,
	}
	e.Orders = append(e.Orders, newOrder)
	matchIds := e.CheckMatchOrder(symbol)
	log.Println(matchIds)
	return &newOrder
}

func (e *Exchange) CheckMatchOrder(symbol string) []uint64 {
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

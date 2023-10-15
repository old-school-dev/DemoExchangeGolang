package api


import (
	"demo-strategy/exchange/demox/order"
	"demo-strategy/constant"
)

func (api *API) CreateBuyOrder(pair string, price float64, amount float64) *order.Order {
	newOrder := order.Order{
		Signer: api.Key,
		Pair: pair,
		Type:   constant.ORDER_LIMIT,
		Side:   constant.SIDE_BUY,
		Amount: amount,
		Price:  price,
		Status: constant.STATUS_PENDING,
	}
	api.CreateNewOrder(newOrder)
	api.CheckMatchOrder(pair)
	return &newOrder
}

func (api *API) CreateSellOrder(pair string, price float64, amount float64) *order.Order {
	newOrder := order.Order{
		Signer: api.Key,
		Pair: pair,
		Type:   constant.ORDER_LIMIT,
		Side:   constant.SIDE_SELL,
		Amount: amount,
		Price:  price,
		Status: constant.STATUS_PENDING,
	}
	api.CreateNewOrder(newOrder)
	api.CheckMatchOrder(pair)
	return &newOrder
}

func (api *API) CreateNewOrder(order order.Order) *order.Order {
	order.Id = api.getNewOrderId()
	api.Orders = append(api.Orders, order)
	// Reduce Wallet Balance
	var (
		amtAdjust float64
		currency string
	)
	pairs := api.GetPairDetails(order.Pair)
	switch order.Side {
	case constant.SIDE_BUY:
		amtAdjust = order.Price * order.Amount
		currency = pairs.Referred
	case constant.SIDE_SELL:
		amtAdjust = order.Amount
		currency = pairs.Currency
	}
	api.AdjustWalletBalance(api.Key, currency, -amtAdjust)
	
	return &order
}

func (api *API) GetOrders() order.Orders {
	var myOrders order.Orders
	for _, order := range api.Orders {
		if order.Signer == api.Key {
			myOrders = append(myOrders, order)
		}
	}
	return api.Orders
}

func (api *API) GetOpenOrders() order.Orders {
	var myOrders order.Orders
	for _, order := range api.Orders {
		if order.Signer == api.Key && order.Status == constant.STATUS_PENDING {
			myOrders = append(myOrders, order)
		}
	}
	return myOrders
}

func (api *API) CancelOrder(id uint64) order.Order {
	for i, order := range api.Orders {
		if order.Signer == api.Key && order.Id == id && order.Status == constant.STATUS_PENDING {
			api.Orders[i].Status = constant.STATUS_CANCEL
			return api.Orders[i]
		}
	}
	return order.Order{}
}

func (api *API) GetOrderDetail(id uint64) order.Order {
	for _, order := range api.Orders {
		if order.Id == id && order.Signer == api.Key {
			return order
		}
	}
	return order.Order{}
}

func (api *API) getNewOrderId() uint64 {
	maxId := uint64(0)
	for _, order := range api.Orders {
		if order.Id > maxId {
			maxId = order.Id
		}
	}
	return maxId
}

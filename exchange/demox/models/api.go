package models

type API struct {
	Key    string
	Secret string
	*Exchange
}

func (api *API) CreateBuyOrder(symbol string, price float64) *Order {
	newOrder := Order{
		Signer: api.Key,
		Symbol: symbol,
		Type:   ORDER_LIMIT,
		Side:   SIDE_BUY,
		Price:  price,
		Status: STATUS_PENDING,
	}
	api.CreateNewOrder(newOrder)
	api.checkMatchOrder(symbol)
	return &newOrder
}

func (api *API) CreateSellOrder(symbol string, price float64) *Order {
	newOrder := Order{
		Signer: api.Key,
		Symbol: symbol,
		Type:   ORDER_LIMIT,
		Side:   SIDE_SELL,
		Price:  price,
		Status: STATUS_PENDING,
	}
	api.CreateNewOrder(newOrder)
	api.checkMatchOrder(symbol)
	return &newOrder
}

func (api *API) CreateNewOrder(order Order) *Order {
	order.Id = api.getNewOrderId()
	api.Orders = append(api.Orders, order)
	return &order
}

func (api *API) GetOrders() Orders {
	return api.Orders
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

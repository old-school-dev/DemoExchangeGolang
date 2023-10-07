package models

type API struct {
	Key    string
	Secret string 
	*Exchange
}

func (api *API) CreateBuyOrder(symbol string, price float64) *Order {
	id := api.getNewId()
	newOrder := Order{
		Id:     id,
		Symbol: symbol,
		Type:   ORDER_LIMIT,
		Side:   SIDE_BUY,
		Price:  price,
		Status: STATUS_PENDING,
	}
	api.Orders = append(api.Orders, newOrder)
	api.checkMatchOrder(symbol)
	return &newOrder
}

func (api *API) CreateSellOrder(symbol string, price float64) *Order {
	id := api.getNewId()
	newOrder := Order{
		Id:     id,
		Symbol: symbol,
		Type:   ORDER_LIMIT,
		Side:   SIDE_SELL,
		Price:  price,
		Status: STATUS_PENDING,
	}
	api.Orders = append(api.Orders, newOrder)
	api.checkMatchOrder(symbol)
	return &newOrder
}

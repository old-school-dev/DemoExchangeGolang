package models

type Order struct {
	Id     uint64
	Type   string
	Side   string
	Price  float64
	Status string
	Symbol string
}

const (
	MAX_PRICE_CHANGE_PERCENT = 1

	SIDE_BUY  = "BUY"
	SIDE_SELL = "SELL"

	ORDER_LIMIT  = "LIMIT"
	ORDER_MARKET = "MARKET"

	STATUS_PENDING = "PENDING"
	STATUS_MATCH   = "MATCH"
	STATUS_CANCEL  = "CANCEL"
)

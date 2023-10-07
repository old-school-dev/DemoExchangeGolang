package models

type Order struct {
	Signer string
	Id     uint64
	Type   string
	Side   string
	Price  float64
	Status string
	Symbol string
}

type Orders []Order
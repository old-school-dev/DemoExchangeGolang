package order

type Order struct {
	Signer string
	Id     uint64
	Type   string
	Side   string
	Price  float64
	Amount float64
	Status string
	Pair string
}

type Orders []Order
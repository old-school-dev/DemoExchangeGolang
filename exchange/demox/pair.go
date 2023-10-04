package demox

import (
	"math"
	"math/rand"
)

type Pair struct {
	Symbol string
	Price  float64
}

func NewPair(symbol string) *Pair {
	startPrice := float64(1000)
	return &Pair{Symbol: symbol, Price: startPrice}
}

func (p *Pair) AdjustPrice() float64 {
	percentChange := (rand.Float64() - 0.5) * 2 * float64(MAX_PRICE_CHANGE_PERCENT)
	p.Price = p.Price + (percentChange / 100 * p.Price)
	p.Price = math.Floor(p.Price*1e2) / 1e2
	return p.Price
}

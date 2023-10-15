package pair

import (
	"demo-strategy/exchange/demox/utils/history"
	"math"
	"math/rand"
	"strings"
)

const MAX_PRICE_CHANGE_PERCENT = 1

type Pair struct {
	Symbol   string
	Price    float64
	Currency string
	Referred  string
	HistoryData *history.History
}

type Pairs []Pair

func NewPair(symbol string) *Pair {
	// startPrice := float64(1000)
	// sliceCurrency := strings.Split(symbol, "/")
	// return &Pair{
	// 	Symbol: symbol, 
	// 	Price: startPrice,
	// 	Currency: sliceCurrency[0],
	// 	Referred: sliceCurrency[1],
	// }
	hist := history.NewHistory()
	startPrice := hist.Next().Open
	sliceCurrency := strings.Split(symbol, "/")
	return &Pair{
		Symbol: symbol, 
		Price: startPrice,
		Currency: sliceCurrency[0],
		Referred: sliceCurrency[1],
		HistoryData: hist,
	}
}

func (p *Pair) adjustPrice() float64 {
	percentChange := (rand.Float64() - 0.5) * 2 * float64(MAX_PRICE_CHANGE_PERCENT)
	p.Price = p.Price + (percentChange / 100 * p.Price)
	p.Price = math.Floor(p.Price*1e2) / 1e2
	return p.Price
}

func (p *Pair) GetPrice() float64 {
	// p.adjustPrice()
	p.Price = p.HistoryData.Next().Open
	return p.Price
}

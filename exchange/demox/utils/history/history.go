package history

import "log"

type History struct {
	history []PriceHistory
	index   int
}

type PriceHistory struct {
	Date  string
	Time  string
	Open  float64
	High  float64
	Low   float64
	Close float64
}

func NewHistory() *History {
	his := LoadHistory()
	return &History{
		history: his,
		index:   0,
	}
}

func (h *History) InsertOne(his PriceHistory) {
	h.history = append(h.history, his)
}

func (h *History) InsertMany(his []PriceHistory) {
	h.history = append(h.history, his...)
}

func (h *History) Next() PriceHistory {
	data := h.history[h.index]
	log.Println(h.index)
	h.index += 1
	return data
}
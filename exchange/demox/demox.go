package demox

func NewExchange(apiKey, apiSecret string) *Exchange{
	Pairs := map[string]*Pair{"BTC": NewPair("BTC")}
	return &Exchange{
		Pairs: Pairs,
		ApiKey: apiKey,
		ApiSecret: apiSecret,
		Orders: []Order{},
	}
}
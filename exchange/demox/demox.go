package demox

import (
	"demo-strategy/exchange/demox/exchange"
	"demo-strategy/exchange/demox/api"
	"demo-strategy/exchange/demox/pair"
	"demo-strategy/exchange/demox/wallet"
)

func NewExchange() *exchange.Exchange {
	pairs := map[string]*pair.Pair{}
	for _, p := range GetPairList() {
		pairs[p] = pair.NewPair(p)
	}
	return &exchange.Exchange{
		Pairs:   pairs,
		Wallets: []wallet.Wallet{},
	}
}

func NewAPI(key, secret string, exchange *exchange.Exchange) *api.API {
	balances := map[string]float64{}
	for _, pair := range GetCurrencyList() {
		balances[pair] = 0
	}
	exchange.CreateWallet(wallet.Wallet{
		ApiKey:   key,
		Balances: balances,
	})
	api := &api.API{
		Key:      key,
		Secret:   secret,
		Exchange: exchange,
	}
	return api
}

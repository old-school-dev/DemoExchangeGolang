package demox

import (
	"demo-strategy/exchange/demox/models"
)

func NewExchange() *models.Exchange {
	pairs := map[string]*models.Pair{}
	for _, pair := range PAIR_LIST {
		pairs[pair] = models.NewPair(pair)
	}
	return &models.Exchange{
		Pairs:   pairs,
		Wallets: []models.Wallet{},
	}
}

func NewAPI(key, secret string, exchange *models.Exchange) *models.API {
	balances := map[string]float64{}
	for _, pair := range PAIR_LIST {
		balances[pair] = 0
	}
	exchange.CreateWallet(models.Wallet{
		ApiKey:   key,
		Balances: balances,
	})
	api := &models.API{
		Key:      key,
		Secret:   secret,
		Exchange: exchange,
	}
	return api
}

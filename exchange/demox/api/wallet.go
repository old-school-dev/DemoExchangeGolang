package api

import "demo-strategy/exchange/demox/wallet"

func (api *API) GetBalances() wallet.Wallet {
	for _, w := range api.Wallets {
		if w.ApiKey == api.Key {
			return w
		}
	}
	return wallet.Wallet{}
}

// func (api *API) GetBalance(currency)

func (api *API) Deposit(symbol string, amount float64) float64 {
	for i, wallet := range api.Wallets {
		if wallet.ApiKey == api.Key {
			api.Wallets[i].Balances[symbol] += amount
			return api.Wallets[i].Balances[symbol]
		}
	}
	return 0.0
}

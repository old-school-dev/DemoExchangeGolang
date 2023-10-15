package wallet

type Wallet struct {
	ApiKey    string
	Balances map[string]float64
}

type Wallets []Wallet
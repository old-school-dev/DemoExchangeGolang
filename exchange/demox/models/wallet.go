package models

type Wallet struct {
	ApiKey    string
	Balances map[string]float64
}

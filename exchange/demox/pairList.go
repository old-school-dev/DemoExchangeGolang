package demox

import "fmt"

type CurrencyList struct {
	Currency []string
	Referred []string
}

var currencyList = CurrencyList{
	Currency: []string{
		"BTC",
		"ETH",
		"BNB",
	},
	Referred: []string{
		"THB",
	},
}

func GetCurrencyList() []string {
	return append(currencyList.Currency, currencyList.Referred...)
}

func GetPairList() []string {
	var pairs []string
	for _, c := range currencyList.Currency {
		for _, r := range currencyList.Referred {
			pairs = append(pairs, fmt.Sprintf("%s/%s",c,r))
		}
	}
	return pairs
}
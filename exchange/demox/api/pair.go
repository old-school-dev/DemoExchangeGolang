package api

import "demo-strategy/exchange/demox/pair"

func (api *API) GetPairDetails(pair string) pair.Pair {
	return *api.Pairs[pair]
}
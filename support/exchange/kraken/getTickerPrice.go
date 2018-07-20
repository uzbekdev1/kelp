package kraken

import (
	"github.com/lightyeario/kelp/model"
	"github.com/lightyeario/kelp/support/exchange/api"
)

// GetTickerPrice impl.
func (k krakenExchange) GetTickerPrice(pairs []model.TradingPair) (map[model.TradingPair]api.Ticker, error) {
	pairsMap, e := model.TradingPairs2Strings(k.assetConverter, k.delimiter, pairs)
	if e != nil {
		return nil, e
	}

	resp, e := k.api.Ticker(values(pairsMap)...)
	if e != nil {
		return nil, e
	}

	priceResult := map[model.TradingPair]api.Ticker{}
	for _, p := range pairs {
		pairTickerInfo := resp.GetPairTickerInfo(pairsMap[p])
		priceResult[p] = api.Ticker{
			AskPrice:  model.MustFromString(pairTickerInfo.Ask[0], k.precision),
			AskVolume: model.MustFromString(pairTickerInfo.Ask[1], k.precision),
			BidPrice:  model.MustFromString(pairTickerInfo.Bid[0], k.precision),
			BidVolume: model.MustFromString(pairTickerInfo.Bid[1], k.precision),
		}
	}

	return priceResult, nil
}

// values gives you the values of a map
// TODO 2 - move to autogenerated generic function
func values(m map[model.TradingPair]string) []string {
	values := []string{}
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
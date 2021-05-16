package yahoo

import "github.com/deifyed/stock-notifier/pkg/stock"

func tickToStockData(tickers []string, ticks []tick) []stock.Data {
	datas := make([]stock.Data, len(ticks))

	for index, tick := range ticks {
		datas[index] = stock.Data{
			Symbol:       tickers[index],
			CurrentPrice: tick.Price,
		}
	}

	return datas
}

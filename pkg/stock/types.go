package stock

type Data struct {
	Symbol       string
	CurrentPrice float64
}

type Client interface {
	FetchCurrentPrice(tickers ...string) ([]Data, error)
}

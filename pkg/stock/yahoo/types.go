package yahoo

import "net/url"

type client struct {
	apiURL url.URL
}

type tick struct {
	YesterdayPrice float64 `json:"regularMarketPreviousClose"`
	Time           int64   `json:"regularMarketTime"`
	Price          float64 `json:"regularMarketPrice"`
}

type quoteResponse struct {
	Result []tick `json:"result"`
}

type response struct {
	QuoteResponse quoteResponse `json:"quoteResponse"`
}

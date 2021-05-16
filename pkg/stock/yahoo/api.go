package yahoo

import (
	"encoding/json"
	"fmt"
	"github.com/deifyed/stock-notifier/pkg/stock"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (c client) FetchCurrentPrice(tickers ...string) ([]stock.Data, error) {
	apiURL := c.apiURL

	query, _ := url.ParseQuery(apiURL.RawQuery)
	query.Add("symbols", strings.Join(tickers, ","))

	apiURL.RawQuery = query.Encode()

	res, err := http.Get(apiURL.String())
	if err != nil {
		return []stock.Data{}, fmt.Errorf("fetching quote: %w", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []stock.Data{}, fmt.Errorf("reading raw quote data: %w", err)
	}

	var response response

	err = json.Unmarshal(data, &response)
	if err != nil {
		return []stock.Data{}, fmt.Errorf("parsing json response: %w", err)
	}

	return tickToStockData(tickers, response.QuoteResponse.Result), nil
}

func NewYahooClient() stock.Client {
	apiURL := url.URL{
		Scheme: "https",
		Host:   "query1.finance.yahoo.com",
		Path:   "/v7/finance/quote",
	}

	query := url.Values{}
	query.Add("lang", "en-US")
	query.Add("region", "US")
	query.Add("corsDomain", "finance.yahoo.com")

	apiURL.RawQuery = query.Encode()

	return &client{apiURL: apiURL}
}

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/deifyed/stock-notifier/pkg/config"
	"github.com/deifyed/stock-notifier/pkg/core"
	"github.com/deifyed/stock-notifier/pkg/notification"
	"github.com/deifyed/stock-notifier/pkg/notification/gotify"
	"github.com/deifyed/stock-notifier/pkg/stock"
	"github.com/deifyed/stock-notifier/pkg/stock/yahoo"
)

func main() {
	cfg := config.LoadConfig(os.Getenv)
	if err := cfg.Validate(); err != nil {
		log.Fatalln(err)
	}

	stockData, err := yahoo.NewYahooClient().FetchCurrentPrice(cfg.Symbols...)
	if err != nil {
		log.Fatalln(fmt.Errorf("fetching quotes: %w", err))
	}

	targetsGetter := generateTargetsGetter(cfg.PriceTargets)
	currentPriceGetter := generateCurrentPriceGetter(stockData)

	stockFactory := core.NewFactory(targetsGetter, currentPriceGetter)

	notificationClient := gotify.NewGotifyClient(cfg.PushServerURL)

	for _, symbol := range cfg.Symbols {
		currentStock := stockFactory.New(symbol)

		if currentStock.TestPriceTargets() {
			err := notificationClient.Notify(notification.Message{
				Title:    "Price target",
				Message:  fmt.Sprintf("%s hit %f", symbol, currentStock.CurrentPrice),
				Priority: 5,
			})
			if err != nil {
				log.Fatalln(fmt.Errorf("sending notification: %w", err))
			}
		}
	}
}

func generateTargetsGetter(rawPriceTargets map[string][]string) func(string) []string {
	return func(symbol string) []string {
		return rawPriceTargets[strings.ToUpper(symbol)]
	}
}

func generateCurrentPriceGetter(datas []stock.Data) func(string) float64 {
	return func(symbol string) float64 {
		for _, item := range datas {
			if strings.EqualFold(item.Symbol, symbol) {
				return item.CurrentPrice
			}
		}

		return -1
	}
}

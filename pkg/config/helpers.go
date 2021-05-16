package config

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func handleSymbols(getter StringGetter, key string) []string {
	rawSymbols := getter(key)
	if rawSymbols == "" {
		log.Fatalln(fmt.Sprintf("getting required env variable %s", key))
	}

	rawSymbols = strings.ReplaceAll(rawSymbols, " ", "")
	rawSymbols = strings.ReplaceAll(rawSymbols, "\n", "")
	rawSymbols = strings.ToUpper(rawSymbols)

	return strings.Split(rawSymbols, ",")
}

func handlePushServerURL(getter StringGetter, key string) url.URL {
	rawURL := getter(key)
	if rawURL == "" {
		log.Fatalln(fmt.Sprintf("getting required env variable %s", key))
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Fatalln(fmt.Sprintf("parsing push server URL %s: %s", rawURL, err.Error()))
	}

	return *parsedURL
}

func handlePriceTargets(getter StringGetter, symbols []string) map[string]string {
	priceTargets := map[string]string{}

	for _, symbol := range symbols {
		rawPriceTargets := getter(fmt.Sprintf("%s_%s", strings.ToUpper(symbol), "TARGETS"))
		if rawPriceTargets == "" {
			log.Fatalln(fmt.Sprintf("missing price targets for %s. Add with %s_TARGETS", symbol, symbol))
		}

		priceTargets[symbol] = strings.ToUpper(rawPriceTargets)
	}

	return priceTargets
}

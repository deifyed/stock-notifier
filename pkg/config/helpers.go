package config

import (
	"fmt"
	"net/url"
	"strings"
)

func handleSymbols(getter StringGetter, key string) []string {
	rawSymbols := getter(key)

	rawSymbols = strings.ToUpper(rawSymbols)

	splitSymbols := strings.Split(rawSymbols, ",")

	return cleanList(splitSymbols)
}

func handlePushServerURL(getter StringGetter, key string) url.URL {
	rawURL := getter(key)
	if rawURL == "" {
		return url.URL{}
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return url.URL{}
	}

	return *parsedURL
}

func handlePriceTargets(getter StringGetter, symbols []string) map[string][]string {
	priceTargets := map[string][]string{}

	for _, symbol := range symbols {
		rawPriceTargets := getter(fmt.Sprintf("%s_%s", strings.ToUpper(symbol), "TARGETS"))

		parts := strings.Split(rawPriceTargets, ",")

		priceTargets[symbol] = cleanList(parts)
	}

	return priceTargets
}

func cleanList(l []string) []string {
	if len(l) == 1 && l[0] == "" {
		return []string{}
	}

	cleanedList := make([]string, len(l))

	for index, item := range l {
		cleanedItem := strings.ReplaceAll(item, " ", "")
		cleanedItem = strings.ReplaceAll(cleanedItem, "\n", "")

		cleanedList[index] = cleanedItem
	}

	return cleanedList
}

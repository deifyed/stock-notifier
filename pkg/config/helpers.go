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

package config

func LoadConfig(getter StringGetter) Config {
	symbols := handleSymbols(getter, "SYMBOLS")
	priceTargets := handlePriceTargets(getter, symbols)

	return Config{
		PushServerURL: handlePushServerURL(getter, "PUSH_SERVER_URL"),
		Symbols:       symbols,
		PriceTargets:  priceTargets,
	}
}

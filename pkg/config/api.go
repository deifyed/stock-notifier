package config

func LoadConfig(getter StringGetter) Config {
	return Config{
		PushServerURL: handlePushServerURL(getter, "PUSH_SERVER_URL"),
		Symbols:       handleSymbols(getter, "SYMBOLS"),
	}
}

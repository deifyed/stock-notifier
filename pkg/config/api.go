package config

import (
	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var (
	tickerRE      = regexp.MustCompile(`[a-zA-Z]+`)
	priceTargetRE = regexp.MustCompile(`([\d]+[+-])`)
)

func LoadConfig(getter StringGetter) Config {
	symbols := handleSymbols(getter, "SYMBOLS")
	priceTargets := handlePriceTargets(getter, symbols)

	return Config{
		PushServerURL: handlePushServerURL(getter, "PUSH_SERVER_URL"),
		Symbols:       symbols,
		PriceTargets:  priceTargets,
	}
}

func (c Config) Validate() (err error) {
	if err = validation.Validate(c.PushServerURL.String(), validation.Required, is.URL); err != nil {
		return err
	}

	if err = validation.Validate(c.Symbols,
		validation.Required,
		validation.Each(validation.Match(tickerRE)),
	); err != nil {
		return err
	}

	if err = validation.Validate(c.PriceTargets,
		validation.Each(validation.Each(validation.Match(priceTargetRE))),
	); err != nil {
		return err
	}

	return nil
}

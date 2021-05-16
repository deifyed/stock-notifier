package core

type (
	TargetsGetter      func(symbol string) (rawTargets string)
	CurrentPriceGetter func(symbol string) (price float64)
)

type Factory struct {
	targetsGetter      TargetsGetter
	currentPriceGetter CurrentPriceGetter
}

func (f Factory) New(symbol string) Stock {
	return Stock{
		Symbol:       symbol,
		PriceTargets: convertTargets(f.targetsGetter(symbol)),
		CurrentPrice: f.currentPriceGetter(symbol),
	}
}

func NewFactory(targetsGetter TargetsGetter, currentPriceGetter CurrentPriceGetter) Factory {
	return Factory{
		targetsGetter:      targetsGetter,
		currentPriceGetter: currentPriceGetter,
	}
}

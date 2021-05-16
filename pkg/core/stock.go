package core

type Stock struct {
	Symbol string

	PriceTargets []PriceTarget

	CurrentPrice float64
}

func (s Stock) TestPriceTargets() bool {
	for _, target := range s.PriceTargets {
		if target.Test(s.CurrentPrice) {
			return true
		}
	}

	return false
}

package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStock(t *testing.T) {
	testCases := []struct {
		name string

		withFactory Factory
		withSymbol  string

		expectStock Stock
	}{
		{
			name: "Should work",

			withFactory: Factory{
				targetsGetter: func(_ string) []string {
					return []string{"100+"}
				},
				currentPriceGetter: func(_ string) float64 {
					return 112
				},
			},
			withSymbol: "JUL",

			expectStock: Stock{
				Symbol: "JUL",
				PriceTargets: []PriceTarget{
					{
						Operator: "+",
						Value:    100,
					},
				},
				CurrentPrice: 112,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			s := tc.withFactory.New(tc.withSymbol)

			assert.Equal(t, tc.expectStock, s)
		})
	}
}

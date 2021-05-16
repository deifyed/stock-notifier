package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStock_TestPriceTargets(t *testing.T) {
	testCases := []struct {
		name string

		withStock Stock
		expectHit bool
	}{
		{
			name: "Should work",

			withStock: Stock{
				Symbol: "",
				PriceTargets: []PriceTarget{
					{
						Operator: "+",
						Value:    10,
					},
				},
				CurrentPrice: 11,
			},
			expectHit: true,
		},
		{
			name: "Should work",

			withStock: Stock{
				Symbol: "",
				PriceTargets: []PriceTarget{
					{
						Operator: "+",
						Value:    10,
					},
				},
				CurrentPrice: 9,
			},
			expectHit: false,
		},
		{
			name: "Should work",

			withStock: Stock{
				Symbol: "",
				PriceTargets: []PriceTarget{
					{
						Operator: "-",
						Value:    5,
					},
				},
				CurrentPrice: 4,
			},
			expectHit: true,
		},
		{
			name: "Should work",

			withStock: Stock{
				Symbol: "",
				PriceTargets: []PriceTarget{
					{
						Operator: "-",
						Value:    5,
					},
				},
				CurrentPrice: 6,
			},
			expectHit: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectHit, tc.withStock.TestPriceTargets())
		})
	}
}

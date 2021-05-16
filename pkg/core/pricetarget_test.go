package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTester(t *testing.T) {
	testCases := []struct {
		name string

		withTarget PriceTarget
		withPrice  float64

		expectResult bool
	}{
		{
			name: "Should return true on + hit",
			withTarget: PriceTarget{
				Operator: "+",
				Value:    10,
			},
			withPrice:    11,
			expectResult: true,
		},
		{
			name: "Should return false on + miss",
			withTarget: PriceTarget{
				Operator: "+",
				Value:    10,
			},
			withPrice:    9,
			expectResult: false,
		},
		{
			name: "Should return true on - hit",
			withTarget: PriceTarget{
				Operator: "-",
				Value:    10,
			},
			withPrice:    9,
			expectResult: true,
		},
		{
			name: "Should return false on - miss",
			withTarget: PriceTarget{
				Operator: "-",
				Value:    10,
			},
			withPrice:    11,
			expectResult: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectResult, tc.withTarget.Test(tc.withPrice))
		})
	}
}

func TestConvertTargets(t *testing.T) {
	testCases := []struct {
		name string

		withRawTargets []string
		expectTargets  []PriceTarget
	}{
		{
			name:           "Should work",
			withRawTargets: []string{"20+", "13-"},
			expectTargets: []PriceTarget{
				{
					Operator: "+",
					Value:    20,
				},
				{
					Operator: "-",
					Value:    13,
				},
			},
		},
		{
			name:           "Should work",
			withRawTargets: []string{"100+", "200+", "300+", "400+"},
			expectTargets: []PriceTarget{
				{
					Operator: "+",
					Value:    100,
				},
				{
					Operator: "+",
					Value:    200,
				},
				{
					Operator: "+",
					Value:    300,
				},
				{
					Operator: "+",
					Value:    400,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			result := convertTargets(tc.withRawTargets)

			assert.Equal(t, tc.expectTargets, result)
		})
	}
}

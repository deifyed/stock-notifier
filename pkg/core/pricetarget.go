package core

import (
	"strconv"
)

type PriceTarget struct {
	Operator string
	Value    float64
}

func (t PriceTarget) Test(price float64) bool {
	switch t.Operator {
	case "+":
		return price > t.Value
	case "-":
		return price < t.Value
	}

	return false
}

// Expected format: "10+,5-, etc"
func convertTargets(rawTargets []string) []PriceTarget {
	targets := make([]PriceTarget, len(rawTargets))

	for index, rawTarget := range rawTargets {
		lastIndex := len(rawTarget) - 1

		operator := rawTarget[lastIndex:]
		rawValue := rawTarget[:lastIndex]

		value, _ := strconv.ParseFloat(rawValue, 64)

		targets[index] = PriceTarget{
			Operator: operator,
			Value:    value,
		}
	}

	return targets
}

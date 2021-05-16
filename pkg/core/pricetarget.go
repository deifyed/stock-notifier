package core

import (
	"strconv"
	"strings"
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

func clean(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\n", "")

	return s
}

// Expected format: "10+,5-, etc"
func convertTargets(rawTargets string) []PriceTarget {
	cleanedRawTargets := clean(rawTargets)

	parts := strings.Split(cleanedRawTargets, ",")

	targets := make([]PriceTarget, len(parts))

	for index, rawTarget := range parts {
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

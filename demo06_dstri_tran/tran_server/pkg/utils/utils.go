package utils

import "math"

func RoundToTwoDecimals(value float64) float64 {
	const precision = 100 // 10^2, for two decimals
	return math.Round(value*precision) / precision
}

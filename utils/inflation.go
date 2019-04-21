package utils

import "math"

// AdjustForInflation takes an arbitrary value and adjusts it for consistent inflation across a
// number of years.
func AdjustForInflation(
	val int64,
	inflationRate float64,
	yearOffset int,
) int64 {

	adjRate := math.Pow(inflationRate, float64(yearOffset))
	adjVal := float64(val) * adjRate
	return int64(adjVal)
}

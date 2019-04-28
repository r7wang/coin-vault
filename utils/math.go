package utils

import "math"

// Inflate takes an arbitrary value and adjusts it at a consistent rate across a number of iterations.
func Inflate(
	val int64,
	rate float64,
	iterations int,
) int64 {

	adjMultiple := math.Pow(1+rate, float64(iterations))
	adjVal := float64(val) * adjMultiple
	return int64(adjVal)
}

// Min returns the smallest of the two int64 values provided. Operations on this type are not
// provided by the default math library.
func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

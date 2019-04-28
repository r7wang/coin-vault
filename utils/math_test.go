package utils_test

import (
	"testing"

	"github.com/r7wang/coin-vault/utils"
	"github.com/stretchr/testify/assert"
)

func Inflate(t *testing.T) {
	tests := []struct {
		name       string
		inputVal   int64
		rate       float64
		iterations int
		expected   int64
	}{
		{
			name:       "with no time and no inflation",
			inputVal:   100,
			rate:       0,
			iterations: 0,
			expected:   100,
		},
		{
			name:       "with zero input and growth",
			inputVal:   0,
			rate:       0.1,
			iterations: 2,
			expected:   0,
		},
		{
			name:       "with time but no inflation",
			inputVal:   100,
			rate:       0,
			iterations: 2,
			expected:   100,
		},
		{
			name:       "with inflation but no time",
			inputVal:   100,
			rate:       0.1,
			iterations: 0,
			expected:   100,
		},
		{
			name:       "with inflation and time",
			inputVal:   100,
			rate:       0.1,
			iterations: 2,
			expected:   121,
		},
		{
			name:       "with inflation and time result rounded down",
			inputVal:   100,
			rate:       0.02,
			iterations: 2,
			expected:   104,
		},
		{
			name:       "with deflation and time",
			inputVal:   100,
			rate:       -0.1,
			iterations: 2,
			expected:   81,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)
			totalTax := utils.Inflate(
				test.inputVal,
				test.rate,
				test.iterations)
			assert.Equal(test.expected, totalTax)
		})
	}
}

func Test_Min(t *testing.T) {
	tests := []struct {
		name     string
		x        int64
		y        int64
		expected int64
	}{
		{
			name:     "with positive values",
			x:        23,
			y:        8,
			expected: 8,
		},
		{
			name:     "with negative values",
			x:        -5,
			y:        -1,
			expected: -5,
		},
		{
			name:     "with equal values",
			x:        3,
			y:        3,
			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)
			totalTax := utils.Min(test.x, test.y)
			assert.Equal(test.expected, totalTax)
		})
	}
}

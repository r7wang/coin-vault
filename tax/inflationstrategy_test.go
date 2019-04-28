package tax_test

import (
	"testing"

	"github.com/r7wang/coin-vault/tax"
	"github.com/stretchr/testify/assert"
)

func Test_GetBrackets(t *testing.T) {
	tests := []struct {
		name               string
		inflationRate      float64
		yearOffset         int
		expectedBaseIncome int64 // base income for the lowest bracket
	}{
		{
			name:               "with no time and no inflation",
			inflationRate:      0,
			yearOffset:         0,
			expectedBaseIncome: 1058200,
		},
		{
			name:               "with time but no inflation",
			inflationRate:      0,
			yearOffset:         2,
			expectedBaseIncome: 1058200,
		},
		{
			name:               "with inflation but no time",
			inflationRate:      0.02,
			yearOffset:         0,
			expectedBaseIncome: 1058200,
		},
		{
			name:               "with inflation and time",
			inflationRate:      0.02,
			yearOffset:         2,
			expectedBaseIncome: 1100951,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)
			strat := tax.NewInflationStrategy(test.inflationRate)
			brackets, err := strat.GetBrackets(test.yearOffset)
			assert.Nil(err)

			definitions := brackets.Definitions
			assert.Equal(len(definitions), 12)
			assert.Equal(test.expectedBaseIncome, definitions[0].BaseIncome)
			assert.Equal(0.5353, brackets.MaximumRate)
		})
	}
}

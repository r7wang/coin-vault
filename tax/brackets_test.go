package tax_test

import (
	"testing"

	"github.com/r7wang/coin-vault/tax"
	"github.com/stretchr/testify/assert"
)

func Test_DefaultBrackets(t *testing.T) {
	testName := "load constant"
	t.Run(testName, func(t *testing.T) {
		assert := assert.New(t)
		brackets, err := tax.DefaultBrackets()
		assert.Nil(err)

		definitions := brackets.Definitions
		assert.Equal(len(definitions), 12)
		assert.Equal(int64(1058200), definitions[0].BaseIncome)
		assert.Equal(0.5353, brackets.MaximumRate)
	})
}

func Test_TotalTax(t *testing.T) {
	tests := []struct {
		name        string
		grossIncome int64
		expected    int64
	}{
		{
			name:        "with zero income",
			grossIncome: 0,
			expected:    0,
		},
		{
			name:        "with negative income",
			grossIncome: -7500,
			expected:    0,
		},
		{
			name:        "with income below zero rate bracket",
			grossIncome: 6300,
			expected:    0,
		},
		{
			name:        "with income equal zero rate bracket",
			grossIncome: 10000,
			expected:    0,
		},
		{
			name:        "with income spanning multiple brackets",
			grossIncome: 25000,
			expected:    2000,
		},
		{
			name:        "with income spanning all brackets",
			grossIncome: 50000,
			expected:    10000,
		},
		{
			name:        "with income crossing highest bracket",
			grossIncome: 64000,
			expected:    17000,
		},
	}

	brackets := makeBrackets()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)
			totalTax := brackets.TotalTax(test.grossIncome)
			assert.Equal(test.expected, totalTax)
		})
	}
}

func makeBrackets() tax.Brackets {
	return tax.Brackets{
		Definitions: []tax.Bracket{
			tax.Bracket{BaseIncome: 10000, TaxRate: 0},
			tax.Bracket{BaseIncome: 20000, TaxRate: 0.1},
			tax.Bracket{BaseIncome: 30000, TaxRate: 0.2},
			tax.Bracket{BaseIncome: 40000, TaxRate: 0.3},
			tax.Bracket{BaseIncome: 50000, TaxRate: 0.4},
		},
		MaximumRate: 0.5,
	}
}

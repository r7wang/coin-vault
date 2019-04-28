package tax

import (
	"github.com/r7wang/coin-vault/utils"
)

// InflationStrategy :
//	* Inflation rate does not change.
//	* Salary increases year-over-year in direct proportion to the rate of inflation.
type InflationStrategy struct {
	inflationRate float64
}

func NewInflationStrategy(inflationRate float64) Strategy {
	return InflationStrategy{
		inflationRate: inflationRate,
	}
}

func (s InflationStrategy) Brackets(yearOffset int) Brackets {
	defaultBrackets, err := DefaultBrackets()
	if err != nil {
		panic("Could not determine default tax brackets.")
	}
	for _, bracket := range defaultBrackets.Definitions {
		bracket.BaseIncome = utils.Inflate(
			bracket.BaseIncome,
			s.inflationRate,
			yearOffset)
	}

	return defaultBrackets
}

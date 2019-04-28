package tax

import (
	"github.com/r7wang/coin-vault/utils"
)

// InflationStrategy :
//	* Inflation rate does not change.
//	* Brackets increase year-over-year in direct proportion to the rate of inflation.
type InflationStrategy struct {
	inflationRate float64
}

func NewInflationStrategy(inflationRate float64) Strategy {
	return InflationStrategy{
		inflationRate: inflationRate,
	}
}

func (s InflationStrategy) GetBrackets(yearOffset int) (Brackets, error) {
	defaultBrackets, err := DefaultBrackets()
	if err != nil {
		return Brackets{}, err
	}

	for idx := range defaultBrackets.Definitions {
		defaultBrackets.Definitions[idx].BaseIncome = utils.Inflate(
			defaultBrackets.Definitions[idx].BaseIncome,
			s.inflationRate,
			yearOffset)
	}

	return defaultBrackets, nil
}

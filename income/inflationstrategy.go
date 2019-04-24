package income

import "math"

// InflationStrategy :
//	* Base salary does not change.
//	* Inflation rate does not change.
//	* Salary increases year-over-year in direct proportion to the rate of inflation.
type InflationStrategy struct {
	baseIncome    int64
	inflationRate float64
}

func NewInflationStrategy(
	baseIncome int64,
	inflationRate float64,
) Strategy {

	return InflationStrategy{
		baseIncome:    baseIncome,
		inflationRate: inflationRate,
	}
}

func (s InflationStrategy) Gross(yearOffset int) int64 {
	adjMultiple := math.Pow(1+s.inflationRate, float64(yearOffset))
	adjIncome := float64(s.baseIncome) * adjMultiple
	return int64(adjIncome)
}

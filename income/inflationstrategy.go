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
) InflationStrategy {

	return InflationStrategy{
		baseIncome:    baseIncome,
		inflationRate: inflationRate,
	}
}

func (s InflationStrategy) Gross(yearOffset int) int64 {
	adjRate := math.Pow(s.inflationRate, float64(yearOffset))
	adjIncome := float64(s.baseIncome) * adjRate
	return int64(adjIncome)
}

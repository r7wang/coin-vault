package tfsa

import (
	"github.com/r7wang/coin-vault/utils"
)

// Calculator performs calculations on TFSA contribution limits.
type Calculator struct {
	inflationRate float64
}

func NewCalculator(inflationRate float64) Calculator {
	return Calculator{
		inflationRate: inflationRate,
	}
}

// ContributionLimit returns the maximum amount that can be contributed to a TFSA, given year
// offset.
func (c Calculator) ContributionLimit(yearOffset int) int64 {
	adjLimit := utils.Inflate(ContributionLimit, c.inflationRate, yearOffset)
	numMultiples := adjLimit / LimitStepSize
	contributionLimit := numMultiples * LimitStepSize
	return contributionLimit
}

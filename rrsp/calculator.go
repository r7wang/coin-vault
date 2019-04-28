package rrsp

import (
	"github.com/r7wang/coin-vault/utils"
)

// Calculator performs calculations on RRSP contribution limits.
type Calculator struct {
	inflationRate float64
}

func NewCalculator(inflationRate float64) Calculator {
	return Calculator{
		inflationRate: inflationRate,
	}
}

// ContributionLimit returns the maximum amount that can be contributed to an RRSP, given gross
// income and year offset.
func (c Calculator) ContributionLimit(
	grossIncome int64,
	yearOffset int,
) int64 {

	maxContributionLimit := utils.Inflate(MaxContributionLimit, c.inflationRate, yearOffset)
	contributionLimit := int64(float64(grossIncome) * MaxContributionPercentage)
	if contributionLimit > maxContributionLimit {
		contributionLimit = maxContributionLimit
	}
	return contributionLimit
}

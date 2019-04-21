package main

import (
	"github.com/r7wang/coin-vault/rrsp"
	"github.com/r7wang/coin-vault/tax"
	"github.com/r7wang/coin-vault/tfsa"
	"github.com/r7wang/coin-vault/utils"
)

// Calculator is a service that handles income taxation.
type Calculator struct {
	taxBrackets   tax.Brackets
	inflationRate float64
}

// NewCalculator ...
func NewCalculator(inflationRate float64) Calculator {
	taxBrackets := tax.NewTaxBrackets(inflationRate)
	return Calculator{
		taxBrackets:   taxBrackets,
		inflationRate: inflationRate,
	}
}

// NetIncome returns net income based on taxation.
func (c Calculator) NetIncome(grossIncome int64, yearOffset int) int64 {
	totalTax := c.taxBrackets.TotalTax(grossIncome, c.inflationRate, yearOffset)
	netIncome := grossIncome - totalTax
	return netIncome
}

// RRSPContributionLimit returns the maximum amount that can be contributed to an RRSP, given income.
func (c Calculator) RRSPContributionLimit(grossIncome int64, yearOffset int) int64 {
	maxContributionLimit := utils.AdjustForInflation(rrsp.MaxContributionLimit, c.inflationRate, yearOffset)
	contributionLimit := int64(float64(grossIncome) * rrsp.MaxContributionPercentage)
	if contributionLimit > maxContributionLimit {
		contributionLimit = maxContributionLimit
	}
	return contributionLimit
}

// TFSAContributionLimit returns the maximum amount that can be contributed to a TFSA.
func (c Calculator) TFSAContributionLimit(yearOffset int) int64 {
	adjLimit := utils.AdjustForInflation(tfsa.ContributionLimit, c.inflationRate, yearOffset)
	numMultiples := adjLimit / tfsa.LimitStepSize
	contributionLimit := numMultiples * tfsa.LimitStepSize
	return contributionLimit
}

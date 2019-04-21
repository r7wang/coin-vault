package main

import (
	"github.com/r7wang/coin-vault/allocation"
	"github.com/r7wang/coin-vault/contribution"
	"github.com/r7wang/coin-vault/income"
)

// Coordinator walks through the main steps of the algorithm.
type Coordinator struct {
	incomeStrategy       income.Strategy
	contributionStrategy contribution.Strategy
	allocationStrategy   allocation.Strategy
	calculator           Calculator
	taxReturns           map[int]int64
}

// Run executes an instance of the algorithm.
func (c Coordinator) Run(
	initialAge int,
	retirementAge int,
	initialBalance Balance) {

	if initialAge > retirementAge {
		panic("Initial age greater than retirement age.")
	}

	currentBalance := initialBalance
	numContributions := retirementAge - initialAge
	for i := 0; i < numContributions; i++ {
		grossIncome := c.incomeStrategy.Gross(i)
		netIncome := c.calculator.NetIncome(grossIncome, i)
		rrspContributionLimit := c.calculator.RRSPContributionLimit(grossIncome, i)
		tfsaContributionLimit := c.calculator.TFSAContributionLimit(i)
		taxReturn := c.getTaxReturn(i)
		totalContribution := c.contributionStrategy.Amount(
			currentBalance,
			netIncome,
			taxReturn,
			rrspContributionLimit,
			tfsaContributionLimit)
		currentBalance = c.allocationStrategy.Allocate(
			currentBalance,
			totalContribution,
			rrspContributionLimit,
			tfsaContributionLimit)

		// TODO: Set tax returns for next year, if any.
	}

	// TODO: Do something with the output, save or print it.
	return
}

func (c Coordinator) getTaxReturn(yearOffset int) int64 {
	if yearOffset == 0 {
		return 0
	}

	taxReturn, ok := c.taxReturns[yearOffset]
	if !ok {
		return 0
	}

	return taxReturn
}

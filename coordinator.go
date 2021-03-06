package main

import (
	"fmt"

	"github.com/r7wang/coin-vault/allocation"
	"github.com/r7wang/coin-vault/contribution"
	"github.com/r7wang/coin-vault/growth"
	"github.com/r7wang/coin-vault/income"
	"github.com/r7wang/coin-vault/rrsp"
	"github.com/r7wang/coin-vault/tax"
	"github.com/r7wang/coin-vault/tfsa"
	"github.com/r7wang/coin-vault/utils"
)

// Coordinator walks through the main steps of the algorithm.
type Coordinator struct {
	growthStrategy       growth.Strategy
	incomeStrategy       income.Strategy
	taxStrategy          tax.Strategy
	contributionStrategy contribution.Strategy
	allocationStrategy   allocation.Strategy
	rrspCalculator       rrsp.Calculator
	tfsaCalculator       tfsa.Calculator
	taxReturns           map[int]int64
}

// Run executes an instance of the algorithm.
func (c Coordinator) Run(
	initialAge int,
	retirementAge int,
	initialBalance utils.Balance) error {

	if initialAge > retirementAge {
		return fmt.Errorf("initial age greater than retirement age")
	}

	currentBalance := initialBalance
	numContributions := retirementAge - initialAge
	for i := 0; i < numContributions; i++ {
		// We are at the start of the year where you turn (initialAge + i).
		//	* You know your salary this year.
		//	* Your initial balance includes all of the contributions you've made so far and there
		//	  are no contributions that haven't been accounted for.
		//	* Throughout the year, you will make enough money to eventually, around end of year,
		//	  contribute a lump sum to your investment accounts.
		//	* Throughout the year, your initial balance will grow according to the growth rate.

		startBalance := currentBalance
		growthRate := c.growthStrategy.Rate(i)
		currentBalance.Grow(growthRate)

		taxBrackets, err := c.taxStrategy.GetBrackets(i)
		if err != nil {
			return err
		}
		grossIncome := c.incomeStrategy.Gross(i)
		regularTax := taxBrackets.TotalTax(grossIncome)
		rrspContributionLimit := c.rrspCalculator.ContributionLimit(grossIncome, i)
		tfsaContributionLimit := c.tfsaCalculator.ContributionLimit(i)
		taxReturn := c.getTaxReturn(i)

		fmt.Println("startBalance", startBalance)
		fmt.Println("currentBalance", currentBalance)
		fmt.Println("taxBrackets", taxBrackets.Definitions)
		fmt.Println("grossIncome", grossIncome)
		fmt.Println("regularTax", regularTax)

		netIncome := grossIncome - regularTax
		totalContribution := c.contributionStrategy.Amount(
			currentBalance,
			netIncome,
			taxReturn,
			rrspContributionLimit,
			tfsaContributionLimit)
		allocation := c.allocationStrategy.Allocate(
			currentBalance,
			totalContribution,
			rrspContributionLimit,
			tfsaContributionLimit)
		currentBalance.AccumulateCash(allocation.Cash)
		currentBalance.AccumulateRRSP(allocation.RRSP, rrspContributionLimit)
		currentBalance.AccumulateTFSA(allocation.TFSA, tfsaContributionLimit)

		// Update future tax returns.
		reducedIncome := grossIncome - allocation.RRSP
		reducedTax := taxBrackets.TotalTax(reducedIncome)
		c.setTaxReturn(i, regularTax-reducedTax)

		/*
			node := EventNode{
				InitialBalance: startBalance,
				NextBalance:    currentBalance,
				GrossIncome:    grossIncome,
				NetIncome:      netIncome,
				TaxReturn:      taxReturn,
				NextTaxReturn:  regularTax - reducedTax,
			}
		*/
		//fmt.Println(node)
	}

	// TODO: Do something with the output, save or print it.
	return nil
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

func (c Coordinator) setTaxReturn(yearOffset int, amount int64) {
	c.taxReturns[yearOffset+1] = amount
}

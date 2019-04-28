package main

import (
	"github.com/r7wang/coin-vault/allocation"
	"github.com/r7wang/coin-vault/contribution"
	"github.com/r7wang/coin-vault/growth"
	"github.com/r7wang/coin-vault/income"
	"github.com/r7wang/coin-vault/rrsp"
	"github.com/r7wang/coin-vault/tax"
	"github.com/r7wang/coin-vault/tfsa"
	"github.com/r7wang/coin-vault/utils"
)

func main() {
	// Get initial situation.
	initialBalance := utils.Balance{
		CashBook:   0,
		CashMarket: 0,
		RRSP:       0,
		RRSPRoom:   0,
		TFSA:       0,
		TFSARoom:   0,
	}

	initialAge := 25
	retirementAge := 55
	growthRate := 0.06
	salary := 90000
	inflationRate := 0.02

	baseSalary := int64(salary * 100)

	coord := Coordinator{
		growthStrategy:       growth.NewStaticRateStrategy(growthRate),
		incomeStrategy:       income.NewInflationStrategy(baseSalary, inflationRate),
		taxStrategy:          tax.NewInflationStrategy(inflationRate),
		contributionStrategy: contribution.RegisteredStrategy{},
		allocationStrategy:   allocation.DefaultStrategy{},
		rrspCalculator:       rrsp.NewCalculator(inflationRate),
		tfsaCalculator:       tfsa.NewCalculator(inflationRate),
		taxReturns:           make(map[int]int64),
	}
	err := coord.Run(initialAge, retirementAge, initialBalance)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/r7wang/coin-vault/allocation"
	"github.com/r7wang/coin-vault/contribution"
	"github.com/r7wang/coin-vault/income"
)

func main() {
	// Get initial situation.
	initialBalance := Balance{
		CashBook:   0,
		CashMarket: 0,
		RRSP:       0,
		RRSPRoom:   0,
		TFSA:       0,
		TFSARoom:   0,
	}

	initialAge := 25
	retirementAge := 55
	salary := 90000
	inflationRate := 1.02

	baseSalary := int64(salary * 100)

	coord := Coordinator{
		incomeStrategy:       income.NewInflationStrategy(baseSalary, inflationRate),
		contributionStrategy: contribution.RegisteredStrategy{},
		allocationStrategy:   allocation.DefaultStrategy{},
		calculator:           NewCalculator(inflationRate),
		taxReturns:           make(map[int]int64),
	}
	coord.Run(initialAge, retirementAge, initialBalance)
}

package main

import "github.com/r7wang/coin-vault/utils"

// EventNode contains information on what was done within a given year.
type EventNode struct {
	InitialBalance utils.Balance
	NextBalance    utils.Balance
	GrossIncome    int64
	NetIncome      int64
	TaxReturn      int64
	NextTaxReturn  int64
}

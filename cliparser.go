package main

type Income struct {
	Amount        int64
	InflationRate float64
}

type CLIParser struct {
	StartingAge        int
	StartingCashBook   int
	StartingCashMarket int
	StartingRRSP       int
	StartingRRSPRoom   int
	StartingTFSA       int
	StartingTFSARoom   int
	WithdrawalAge      int
	WithdrawalAmount   int
	WithdrawalStrategy int
	InflationRate      float64
	YearlyCash         int

	// Withdrawal strategies include:
	//	* Start draining RRSP/RRIF, don't touch TFSA: intent is to frontload tax burden and maximize TFSA compounding.
	//	* Minimum draining RRSP/RRIF, touch TFSA to compensate for rest: intent is to defer tax rate burden until later.
	//	* Mixed drain RRSP/RRIF/TFSA, use RRSP/TFSA in ratios for remaining (90/10 <--> 10/90).

	// RRSP contribution strategy:
	//	* Contribute until you will take it out.
	//	* Stop contributing x years until you will take it out, and instead move to cash account.

	// Cash account capital gains trigger:
	//	* Full cap trigger every year.
	//	* No cap trigger until money needed.
}

func NewCLIParser() (*CLIParser, error) {
	return &CLIParser{}, nil
}

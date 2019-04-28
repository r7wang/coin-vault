package utils

// Balance represents monetary values of personal assets in cents.
type Balance struct {
	CashBook   int64
	CashMarket int64
	RRSP       int64
	RRSPRoom   int64
	TFSA       int64
	TFSARoom   int64
}

// Grow :
// Adjusts the balance based on a provided growth rate.
func (b *Balance) Grow(growthRate float64) {
	growthMultiple := 1 + growthRate
	b.CashMarket = int64(float64(b.CashMarket) * growthMultiple)
	b.RRSP = int64(float64(b.RRSP) * growthMultiple)
	b.TFSA = int64(float64(b.TFSA) * growthMultiple)
}

// AccumulateCash :
// Adjusts the cash balance given an amount.
func (b *Balance) AccumulateCash(amount int64) {
	b.CashBook += amount
	b.CashMarket += amount
}

// AccumulateRRSP :
// Adjusts the RRSP balance and contribution limit given an amount.
func (b *Balance) AccumulateRRSP(amount int64, contributionLimit int64) {
	b.RRSP += amount
	b.RRSPRoom += contributionLimit - amount
}

// AccumulateTFSA :
// Adjusts the TFSA balance and contribution limit given an amount.
func (b *Balance) AccumulateTFSA(amount int64, contributionLimit int64) {
	b.TFSA += amount
	b.TFSARoom += contributionLimit - amount
}

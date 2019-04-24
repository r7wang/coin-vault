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
func (b Balance) Grow(growthRate float64) {
	growthMultiple := 1 + growthRate
	b.CashMarket = int64(float64(b.CashMarket) * growthMultiple)
	b.RRSP = int64(float64(b.RRSP) * growthMultiple)
	b.TFSA = int64(float64(b.TFSA) * growthMultiple)
}

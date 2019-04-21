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

package tax

// Bracket represents a tax bracket.
type Bracket struct {
	// Income from the last bracket up to this value is taxed at the tax rate below.
	BaseIncome int64
	// The tax rate, applied only to this specific bracket.
	TaxRate float64
}

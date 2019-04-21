package income

// Strategy defines an income strategy.
type Strategy interface {
	Gross(yearOffset int) int64
}

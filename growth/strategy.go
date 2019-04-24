package growth

// Strategy defines a growth strategy.
type Strategy interface {
	Rate(yearOffset int) float64
}

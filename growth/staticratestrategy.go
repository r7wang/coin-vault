package growth

// StaticRateStrategy :
// Growth rate is set to a fixed percentage and never changes.
type StaticRateStrategy struct {
	growthRate float64
}

func NewStaticRateStrategy(growthRate float64) Strategy {
	return StaticRateStrategy{
		growthRate: growthRate,
	}
}

func (s StaticRateStrategy) Rate(yearOffset int) float64 {
	return s.growthRate
}

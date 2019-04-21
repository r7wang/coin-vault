package tfsa

const (
	// ContributionLimit :
	// The maximum amount that can be contributed to a TFSA.
	ContributionLimit = int64(600000)

	// LimitStepSize :
	// Changes to the contribution limit must change by some integer multiple of this size.
	LimitStepSize = int64(50000)
)

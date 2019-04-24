package allocation

import "github.com/r7wang/coin-vault/utils"

// Strategy defines a contribution strategy.
type Strategy interface {
	Allocate(
		currentBalance utils.Balance,
		totalContribution int64,
		rrspContributionLimit int64,
		tfsaContributionLimit int64) Allocation
}

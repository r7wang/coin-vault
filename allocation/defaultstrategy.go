package allocation

import "github.com/r7wang/coin-vault/utils"

type DefaultStrategy struct{}

func (s DefaultStrategy) Allocate(
	currentBalance utils.Balance,
	totalContribution int64,
	rrspContributionLimit int64,
	tfsaContributionLimit int64,
) Allocation {

	remainingContribution := totalContribution

	// TFSA (1).
	tfsaContribution := utils.Min(remainingContribution, tfsaContributionLimit)
	remainingContribution -= tfsaContribution

	// RRSP (2).
	rrspContribution := utils.Min(remainingContribution, rrspContributionLimit)
	remainingContribution -= rrspContribution

	// Cash (3).
	cashContribution := remainingContribution

	return Allocation{
		Cash: cashContribution,
		RRSP: rrspContribution,
		TFSA: tfsaContribution,
	}
}

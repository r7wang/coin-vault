package allocation

import "github.com/r7wang/coin-vault/utils"

type DefaultStrategy struct{}

func (s DefaultStrategy) Allocate(
	currentBalance utils.Balance,
	totalContribution int64,
	rrspContributionLimit int64,
	tfsaContributionLimit int64,
) utils.Balance {

	return utils.Balance{}
}

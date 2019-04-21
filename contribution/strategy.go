package contribution

import "github.com/r7wang/coin-vault/utils"

// Strategy defines a contribution strategy.
type Strategy interface {
	Amount(
		currentBalance utils.Balance,
		netIncome int64,
		taxReturn int64,
		rrspContributionLimit int64,
		tfsaContributionLimit int64) int64
}

package contribution

import "github.com/r7wang/coin-vault/utils"

// RegisteredStrategy attempts to contribute whatever is necessary to max out the RRSP and TFSA for
// the current calendar year. Any extra cash, for instance, from last year's tax return but paid
// out this year, is considered a bonus and will be used for contribution. Additional RRSP and TFSA
// contribution room from previous years is ignored.
type RegisteredStrategy struct{}

// Amount returns the amount of money that will be used to contribute to investment accounts.
func (s RegisteredStrategy) Amount(
	currentBalance utils.Balance,
	netIncome int64,
	taxReturn int64,
	rrspContributionLimit int64,
	tfsaContributionLimit int64,
) int64 {

	contributionLimit := rrspContributionLimit + tfsaContributionLimit
	if netIncome < contributionLimit {
		// Even if income is insufficient, we can still try to achieve our objective with extra
		// money from the tax return.
		return netIncome + taxReturn
	}

	return contributionLimit + taxReturn
}

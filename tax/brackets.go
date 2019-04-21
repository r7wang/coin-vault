package tax

import (
	"math"

	"github.com/r7wang/coin-vault/utils"
)

type Brackets []Bracket

// NewTaxBrackets ...
func NewTaxBrackets(inflationRate float64) Brackets {
	return Brackets{
		Bracket{1058200, 0},
		Bracket{1206900, 0.0505},
		Bracket{4390600, 0.2005},
		Bracket{4763000, 0.2405},
		Bracket{7731300, 0.2965},
		Bracket{8781300, 0.3148},
		Bracket{9110100, 0.3389},
		Bracket{9525900, 0.3791},
		Bracket{14766700, 0.4341},
		Bracket{15000000, 0.4641},
		Bracket{21037100, 0.4797},
		Bracket{22000000, 0.5197},
		Bracket{math.MaxInt64, 0.5353}}
}

func (b Brackets) TotalTax(grossIncome int64, inflationRate float64, yearOffset int) int64 {
	accumTax := int64(0)
	for bracketIdx, bracket := range b {
		curBracketIncome := b.getBracketIncome(bracket, inflationRate, yearOffset)
		lastBracketIncome := b.getLastBracketIncome(bracketIdx, inflationRate, yearOffset)
		var bracketTax float64
		if grossIncome >= curBracketIncome {
			bracketTax = bracket.TaxRate * float64(curBracketIncome-lastBracketIncome)
		} else {
			bracketTax = bracket.TaxRate * float64(grossIncome-lastBracketIncome)
		}
		accumTax += int64(bracketTax)
	}
	return accumTax
}

func (b Brackets) getBracketIncome(bracket Bracket, inflationRate float64, yearOffset int) int64 {
	adjIncome := utils.AdjustForInflation(bracket.BaseIncome, inflationRate, yearOffset)
	return adjIncome
}

func (b Brackets) getLastBracketIncome(bracketIdx int, inflationRate float64, yearOffset int) int64 {
	if bracketIdx == 0 {
		return 0
	}

	baseIncome := b[bracketIdx-1].BaseIncome
	adjIncome := utils.AdjustForInflation(baseIncome, inflationRate, yearOffset)
	return adjIncome
}

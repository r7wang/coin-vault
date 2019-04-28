package tax

import (
	"encoding/json"
)

type Brackets struct {
	Definitions []Bracket
	MaximumRate float64
}

// DefaultBrackets returns the default tax brackets for the current year.
func DefaultBrackets() (Brackets, error) {
	var brackets Brackets
	if err := json.Unmarshal([]byte(defaultBrackets), &brackets); err != nil {
		return Brackets{}, err
	}

	return brackets, nil
}

// TotalTax returns the total amount of tax that would be paid for a given gross income and a set
// of tax brackets.
func (b Brackets) TotalTax(grossIncome int64) int64 {
	accumTax := int64(0)
	for bracketIdx, bracket := range b.Definitions {
		curBracketIncome := bracket.BaseIncome
		lastBracketIncome := b.getLastBracketIncome(bracketIdx)
		var bracketTax float64
		if grossIncome >= curBracketIncome {
			bracketTax = bracket.TaxRate * float64(curBracketIncome-lastBracketIncome)
		} else if grossIncome >= lastBracketIncome {
			bracketTax = bracket.TaxRate * float64(grossIncome-lastBracketIncome)
		} else {
			break
		}
		accumTax += int64(bracketTax)
	}

	numBrackets := len(b.Definitions)
	finalBracketIncome := b.Definitions[numBrackets-1].BaseIncome
	if grossIncome > finalBracketIncome {
		bracketTax := b.MaximumRate * float64(grossIncome-finalBracketIncome)
		accumTax += int64(bracketTax)
	}

	return accumTax
}

func (b Brackets) getLastBracketIncome(bracketIdx int) int64 {
	if bracketIdx == 0 {
		return 0
	}

	bracketIncome := b.Definitions[bracketIdx-1].BaseIncome
	return bracketIncome
}

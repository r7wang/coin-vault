package tax

const (
	// DefaultBrackets :
	// Tax brackets associated with the current year.
	defaultBrackets = `
		{
			"definitions": [
				{ "baseIncome": 1058200, "taxRate": 0 },
				{ "baseIncome": 1206900, "taxRate": 0.0505 },
				{ "baseIncome": 4390600, "taxRate": 0.2005 },
				{ "baseIncome": 4763000, "taxRate": 0.2405 },
				{ "baseIncome": 7731300, "taxRate": 0.2965 },
				{ "baseIncome": 8781300, "taxRate": 0.3148 },
				{ "baseIncome": 9110100, "taxRate": 0.3389 },
				{ "baseIncome": 9525900, "taxRate": 0.3791 },
				{ "baseIncome": 14766700, "taxRate": 0.4341 },
				{ "baseIncome": 15000000, "taxRate": 0.4641 },
				{ "baseIncome": 21037100, "taxRate": 0.4797 },
				{ "baseIncome": 22000000, "taxRate": 0.5197 }
			],
			"maximumRate": 0.5353
		}
		`
)

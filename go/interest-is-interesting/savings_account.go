package savings

// InterestRate calculates the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return -3.213
	} else if balance < 1000 {
		return 0.5
	} else if balance < 5000 {
		return 1.621
	} else {
		return 2.475
	}
}

// InterestRate calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	if balance < 0 {
		return float64(InterestRate(balance)) / 100 * balance * -1
	}
	return float64(InterestRate(balance)) / 100 * balance
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		years += 1
	}
	return years
}

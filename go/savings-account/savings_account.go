package savings

const FixedInterestRate = .05

// GetFixedInterestRate returns the FixedInterestRate constant
func GetFixedInterestRate() float32 {
	return FixedInterestRate
}

const DaysPerYear = 365

// GetDaysPerYear returns the DaysPerYear constant
func GetDaysPerYear() int {
	return DaysPerYear
}

// Jan-Dec have values of 1-12
const (
	Undefined int = iota
	Jan
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)

// GetMonth returns the value for the given month
func GetMonth(month int) int {
	switch month {
	case Jan:
		return Jan
	case Feb:
		return Feb
	case Mar:
		return Mar
	case Apr:
		return Apr
	case May:
		return May
	case Jun:
		return Jun
	case Jul:
		return Jul
	case Aug:
		return Aug
	case Sep:
		return Sep
	case Oct:
		return Oct
	case Nov:
		return Nov
	case Dec:
		return Dec
	default:
		return Undefined
	}
}

type AccNo string

func GetAccountNumber() AccNo {
	accounts := map[int]AccNo{
		1: "XF348IJ",
	}
	return accounts[1]
}

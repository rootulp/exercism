package perfect

import "errors"

type Classification string

const ClassificationAbundant Classification = "ClassificationAbundant"
const ClassificationDeficient Classification = "ClassificationDeficient"
const ClassificationPerfect Classification = "ClassificationPerfect"

var ErrOnlyPositive = errors.New("error only positive numbers can be classified")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	}
	sum := aliquotSum(n)

	if sum < n {
		return ClassificationDeficient, nil
	} else if sum == n {
		return ClassificationPerfect, nil
	} else if sum > n {
		return ClassificationAbundant, nil
	}
	return "", errors.New("unreachable")
}

func aliquotSum(n int64) (sum int64) {
	factors := getFactors(n)
	for _, factor := range factors {
		sum += factor
	}
	return sum
}

func getFactors(n int64) (factors []int64) {
	for candidate := int64(1); candidate < n; candidate++ {
		if n%candidate == 0 {
			factors = append(factors, candidate)
		}
	}
	return factors
}

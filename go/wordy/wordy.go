package wordy

import (
	"fmt"
	"strconv"
	"strings"
)

// Operators
const plus = "plus"
const minus = "minus"
const multipliedBy = "multiplied by"
const dividedBy = "divided by"
const multiplied = "multiplied"
const divided = "divided"

func Answer(question string) (result int, ok bool) {
	if trim(question) == "" {
		return 0, false
	}

	operands, operators, err := getTokens(question)
	if err != nil {
		return 0, false
	}

	result, err = compute(operands, operators)
	if err != nil {
		return 0, false
	}
	return result, true
}

func compute(operands []int, operators []string) (result int, err error) {
	if len(operands) == 1 && len(operators) == 0 {
		return operands[0], nil
	}
	if len(operators)+1 != len(operands) {
		return 0, fmt.Errorf("there should be one more operator %v than operands %v", operators, operands)
	}

	result = operands[0]
	operands = operands[1:]
	for len(operators) > 0 {
		if len(operators) != len(operands) {
			return 0, fmt.Errorf("uneven number of operators %v and operands %v", operators, operands)
		}
		result, err = execute(operators[0], result, operands[0])
		if err != nil {
			return 0, err
		}
		operators = operators[1:]
		operands = operands[1:]

	}
	return result, nil
}

func execute(operator string, a int, b int) (result int, e error) {
	switch operator {
	case plus:
		return a + b, nil
	case minus:
		return a - b, nil
	case multiplied:
		return a * b, nil
	case divided:
		return a / b, nil
	default:
		return 0, fmt.Errorf("invalid operator %v", operator)
	}
}

func getTokens(question string) (operands []int, operators []string, e error) {
	s := trim(question)

	// Replace two word operators with one word equivalant so we can split on whitespace
	s = strings.ReplaceAll(s, multipliedBy, multiplied)
	s = strings.ReplaceAll(s, dividedBy, divided)
	for i, v := range strings.Fields(s) {
		if i%2 == 1 {
			// Parse operator
			if !isOperator(v) {
				return operands, operators, fmt.Errorf("%v is not an operator", v)
			}
			operators = append(operators, v)
		} else {
			// Parse operand
			operand, err := strconv.Atoi(v)
			if err != nil {
				return operands, operators, fmt.Errorf("%v is not an operand", v)
			}
			operands = append(operands, operand)
		}
	}
	fmt.Printf("operands=%v, operators=%v, err=%v\n", operands, operators, e)
	return operands, operators, nil
}

func isOperator(s string) bool {
	return s == plus || s == minus || s == multiplied || s == divided
}

// trim removes the generic prefix and suffix from a question
func trim(question string) (trimmed string) {
	trimmed = strings.TrimPrefix(question, "What is")
	trimmed = strings.TrimSuffix(trimmed, "?")
	return trimmed
}

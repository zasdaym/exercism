package wordy

import (
	"strconv"
	"strings"
)

type operatorFunc func(a, b int) int

func operatorAddition(a, b int) int {
	return a + b
}

func operatorSubstraction(a, b int) int {
	return a - b
}

func operatorMultiplication(a, b int) int {
	return a * b
}

func operatorDivision(a, b int) int {
	return a / b
}

func Answer(question string) (int, bool) {
	var numbers []int
	var operators []operatorFunc

	question = strings.ReplaceAll(question, "What is ", "")
	question = strings.ReplaceAll(question, "by ", "")
	question = strings.ReplaceAll(question, "?", "")

	tokens := strings.Split(question, " ")
	expectNumber := true

	for _, token := range tokens {
		if expectNumber {
			n, err := strconv.Atoi(token)
			if err != nil {
				return 0, false
			}
			numbers = append(numbers, n)
			expectNumber = false
			continue
		}

		switch token {
		case "plus":
			operators = append(operators, operatorAddition)
		case "minus":
			operators = append(operators, operatorSubstraction)
		case "multiplied":
			operators = append(operators, operatorMultiplication)
		case "divided":
			operators = append(operators, operatorDivision)
		default:
			return 0, false
		}

		expectNumber = true
	}

	// No number found
	if len(numbers) == 0 {
		return 0, false
	}

	// Number of operands should be more than the number of operators
	if len(numbers)-len(operators) != 1 {
		return 0, false
	}

	// Single number
	if len(numbers) == 1 && len(operators) == 0 {
		return numbers[0], true
	}

	result, _ := calculate(numbers, operators)
	return result[0], true
}

func calculate(numbers []int, operators []operatorFunc) ([]int, []operatorFunc) {
	if len(numbers) == 1 {
		return numbers, nil
	}

	a, b := numbers[0], numbers[1]
	op := operators[0]

	newNumbers := []int{op(a, b)}
	newNumbers = append(newNumbers, numbers[2:]...)
	newOperators := operators[1:]

	return calculate(newNumbers, newOperators)
}

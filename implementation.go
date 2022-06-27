package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

var errorMessages = map[string]string {
	"wrongInputForm": "Невірна форма запису виразу! Має використовуватись постфіксна форма!",
	"wrongOperation": "Виконати дану операцію над аргументами неможливо:",
}

// TODO: document this function.
// PrefixToPostfix converts
func CountPostfix(input string) (string, error) {
	args := strings.Split(input, " ")
	if len(args) == 1 {
		number := args[0]
		_, err := strconv.Atoi(number)
		if err != nil {
			return "0", fmt.Errorf(errorMessages["wrongInputForm"])
		}
		return number, nil
	}

	numbers := make([]int, 0)
	for _, val := range args {
		num, err := strconv.Atoi(val)
		if err != nil {
			res, err := operationArgument(numbers, val)
			if err != nil {
				return "0", err
			}
			numbers = res
		} else {
			numbers = numberArgument(numbers, num)
		}
	}

	if len(numbers) > 1 {
		return "0", fmt.Errorf(errorMessages["wrongInputForm"])
	}

	res := strconv.Itoa(numbers[0])
	return res, nil
}

func numberArgument(numbers []int, num int) []int {
	numbers = append(numbers, num)
	return numbers
}

func operationArgument(numbers []int, operation string) ([]int, error) {
	if len(numbers) < 2 {
		return make([]int, 0), fmt.Errorf(errorMessages["wrongInputForm"])
	}

	a := numbers[0]
	b := numbers[1]
	res, err := performOperation(a, b, operation)
	if err != nil {
		return make([]int, 0), err
	}

	if len(numbers) == 2 {
		numbers[0] = res
		numbers = numbers[:1]
	} else {
		numbers[0] = res 
		tmp := numbers[2:]
		numbers = numbers[:1]
		numbers = append(numbers, tmp...)
	}
	return numbers, nil
}

func performOperation(a, b int, operation string) (int, error) {
	switch operation {
	case "+":
		a += b
	case "-":
		a -= b
	case "*":
		a *= b
	case "/":
		a /= b
	default:
		return 0, fmt.Errorf("%s %s!", errorMessages["wrongOperation"], operation)
	}
	return a, nil
}

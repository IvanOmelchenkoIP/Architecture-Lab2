package lab2

import (
	"fmt"
	"strconv"
	"strings"
	"math"
)

var errorMessages = map[string]string {
	"wrongInputForm": "Невірна форма запису виразу! Має використовуватись постфіксна форма!",
	"wrongOperation": "Виконати дану операцію над аргументами неможливо:",
}

type Numbers struct {
	stack []int
}

func (numbers *Numbers) push(num int) {
	numbers.stack = append(numbers.stack, num)
}

func (numbers *Numbers) pop() int {
	ind := numbers.length() - 1
	num := numbers.stack[ind]
	numbers.stack = numbers.stack[:ind]
	return num
}

func (numbers *Numbers) length() int {
	return len(numbers.stack)
}

// TODO: document this function.
// PrefixToPostfix converts
func CountPostfix(input string) (string, error) {
	args := strings.Split(input, " ")
	if len(args) == 1 {
		number := args[0]
		_, err := strconv.Atoi(number)
		if err != nil {
			return "", fmt.Errorf(errorMessages["wrongInputForm"])
		}
		return number, nil
	}

	numbers := Numbers{ stack: make([]int, 0) }
	for _, val := range args {
		num, err := strconv.Atoi(val)
		if err != nil {
			res, err := processOperArg(numbers, val)
			if err != nil {
				return "", err
			}
			numbers = res
		} else {
			numbers = processNumArg(numbers, num)
		}
	}

	if numbers.length() > 1 {
		return "", fmt.Errorf(errorMessages["wrongInputForm"])
	}

	res := strconv.Itoa(numbers.pop())
	return res, nil
}

func processNumArg(numbers Numbers, num int) Numbers {
	numbers.push(num)
	return numbers
}

func processOperArg(numbers Numbers, operation string) (Numbers, error) {
	if numbers.length() < 2 {
		return Numbers{ stack: make([]int, 0) }, fmt.Errorf(errorMessages["wrongInputForm"])
	}

	a := numbers.pop()
	b := numbers.pop()
	res, err := performOperation(a, b, operation)
	if err != nil {
		return Numbers{ stack: make([]int, 0) }, err
	}
	numbers.push(res)
	return numbers, nil
}

func performOperation(a, b int, operation string) (int, error) {
	switch operation {
	case "+":
		b += a
	case "-":
		b -= a
	case "*":
		b *= a
	case "/":
		b /= a
	case "^":
		b = int(math.Pow(float64(b), float64(a)))
	default:
		return 0, fmt.Errorf("%s %s!", errorMessages["wrongOperation"], operation)
	}
	return b, nil
}

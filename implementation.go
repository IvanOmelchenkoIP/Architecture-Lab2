package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var errorMessages = map[string]string{
	"wrongInputForm": "Невірна форма запису виразу! Має використовуватись постфіксна форма!",
	"wrongOperation": "Виконати дану операцію над аргументами неможливо:",
}

type numbers struct {
	stack []int
}

func (nums *numbers) push(num int) {
	nums.stack = append(nums.stack, num)
}

func (nums *numbers) pop() int {
	ind := nums.length() - 1
	num := nums.stack[ind]
	nums.stack = nums.stack[:ind]
	return num
}

func (nums *numbers) length() int {
	return len(nums.stack)
}

// CountPostfix is a function that returns the result of a mathematical expression in postfix notation.
// Work principles: CountPostfix splits its string argument it into an array by spaces (" ").
// If the resulting array has single argument and it is a number - returns it, else - returns error.
// If there are more than one elements in the array, then each element is evaluated:
// if it is a number, it is pushed to a stack; if it is not, the program assumes that it is an operator
// and will try to count result of two last elements of the stack. In the end, CountPostfix should return
// a string, which is the result of the evaluation of the entire expression, or an error.
// Error types: "Wrong form of expression! Postfix notation must be used!" - occurs when the only element of
// the split string array is not a number, when there is overusage or underusage of operators.
// "Operation can not be performed on the arguments: <operation type>" - occurs when it is requested to
// perform an operationon two arguments which is not possible
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

	nums := numbers{stack: make([]int, 0)}
	for _, val := range args {
		num, err := strconv.Atoi(val)
		if err != nil {
			res, err := processOperArg(nums, val)
			if err != nil {
				return "", err
			}
			nums = res
		} else {
			nums = processNumArg(nums, num)
		}
	}

	if nums.length() > 1 {
		return "", fmt.Errorf(errorMessages["wrongInputForm"])
	}

	res := strconv.Itoa(nums.pop())
	return res, nil
}

func processNumArg(nums numbers, num int) numbers {
	nums.push(num)
	return nums
}

func processOperArg(nums numbers, operation string) (numbers, error) {
	if nums.length() < 2 {
		return numbers{stack: make([]int, 0)}, fmt.Errorf(errorMessages["wrongInputForm"])
	}

	a := nums.pop()
	b := nums.pop()
	res, err := performOperation(a, b, operation)
	if err != nil {
		return numbers{stack: make([]int, 0)}, err
	}
	nums.push(res)
	return nums, nil
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

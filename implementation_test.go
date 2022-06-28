package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToPostfix(t *testing.T) {
	res, err := CountPostfix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 2 - 3 * 5 +", res)
	}
}

func ExampleCountPostfix() {
	res, _ := CountPostfix("2 2 + 4 2 * 3 - +")
	fmt.Println(res)

	// Output:
	// 9
}
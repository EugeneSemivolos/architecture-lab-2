package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name       string
	expression string
	wanted     int
	err        error
}

var Tests = []Test{
	{
		name:       "Plus",
		expression: "+ 5 4",
		wanted:     9,
		err:        nil,
	},
	{
		name:       "Minus",
		expression: "- 22 12",
		wanted:     10,
		err:        nil,
	},
	{
		name:       "Multiply",
		expression: "* 8 7",
		wanted:     56,
		err:        nil,
	},
	{
		name:       "Divide",
		expression: "/ 72 9",
		wanted:     8,
		err:        nil,
	},
	{
		name:       "Pow",
		expression: "^ 3 4",
		wanted:     81,
		err:        nil,
	},
	{
		name:       "Complex",
		expression: "^ 2 * 2 / + 3 15 6",
		wanted:     64,
		err:        nil,
	},
	{
		name:       "With spaces",
		expression: "  ^ 2* 2    / +    3 15  6  ",
		wanted:     64,
		err:        nil,
	},
	{
		name:       "NaN",
		expression: "+ b 1 2 3",
		wanted:     -1,
		err:        fmt.Errorf("Error. Problem with opertator"),
	},
	{
		name:       "Divide by 0",
		expression: "/ 72 0",
		wanted:     -1,
		err:        fmt.Errorf("Error. Divide by zero"),
	},
	{
		name:       "Just arguments",
		expression: "1 2 3",
		wanted:     -1,
		err:        fmt.Errorf("Error. Missing arguments or many operators"),
	},
	{
		name:       "Too many operators",
		expression: "+ ^ 2 3",
		wanted:     -1,
		err:        fmt.Errorf("Error. Missing arguments or many operators"),
	},
}

func TestCalculatePrefix(t *testing.T) {
	for _, tc := range Tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := CalculatePrefix(tc.expression)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.wanted, res)
		})
	}
}

func ExampleCalculatePrefix() {
	res, _ := CalculatePrefix("+ 5 * - 4 2 3")
	fmt.Println(res)

	// Output:
	// 11
}

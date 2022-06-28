package lab2

import (
	"fmt"
	"strings"
	"bytes"

	gocheck "gopkg.in/check.v1"
)

type MySuite struct{}

var _ = gocheck.Suite(&MySuite{})

func (s *MySuite) TestComputeOperandComplexity3_1(c *gocheck.C) {
	var input = "4 9 3 / -" // 4 - 9 / 3 - 4 = 1
	var expected = "1"
	var out bytes.Buffer
	handler := &ComputeHandler{
		Input: strings.NewReader(input),
		Output: &out,
	}
	err := handler.Compute()
	if err != nil {
		c.Fail()
	}
	res := strings.Trim(out.String(), "\n")
	c.Assert(res, gocheck.Equals, expected)
}

func (s *MySuite) TestComputeOperandComplexity5_1(c *gocheck.C) {
	var input = "4 1 - 4 3 * 11 - ^" // (4 - 1)^(4 * 3 - 11)
	var expected = "3"
	var out bytes.Buffer
	handler := &ComputeHandler{
		Input: strings.NewReader(input),
		Output: &out,
	}
	err := handler.Compute()
	if err != nil {
		c.Fail()
	}
	res := strings.Trim(out.String(), "\n")
	c.Assert(res, gocheck.Equals, expected)
}


func (s *MySuite) TestComputefixWrongInput_1(c *gocheck.C) {
	var input = "b" // there is only 1 argument and it is not a number, so it can not be result
	var expectedErr = "Невірна форма запису виразу! Має використовуватись постфіксна форма!"
	var out bytes.Buffer
	handler := &ComputeHandler{
		Input: strings.NewReader(input),
		Output: &out,
	}
	err := handler.Compute()
	c.Assert(err, gocheck.ErrorMatches, expectedErr)
}

func (s *MySuite) TestComputeWrongInput_2(c *gocheck.C) {
	var input = "4 2 1 + - + -" // there is not enough arguments to count, so it counts as a postfix notation violation
	var expectedErr = "Невірна форма запису виразу! Має використовуватись постфіксна форма!"
	var out bytes.Buffer
	handler := &ComputeHandler{
		Input: strings.NewReader(input),
		Output: &out,
	}
	err := handler.Compute()
	c.Assert(err, gocheck.ErrorMatches, expectedErr)
}

func (s *MySuite) TestComputeWrongInput_3(c *gocheck.C) {
	var input = "12 3 2 4 / ^" // there are not enough operators to process all the operands, so the notation is violated
	var expectedErr = "Невірна форма запису виразу! Має використовуватись постфіксна форма!"
	var out bytes.Buffer
	handler := &ComputeHandler{
		Input: strings.NewReader(input),
		Output: &out,
	}
	err := handler.Compute()
	c.Assert(err, gocheck.ErrorMatches, expectedErr)
}

func (s *MySuite) TestComputeWrongInput_4(c *gocheck.C) {
	var input = "34 b" // though the operator is wrong, there are not enough arguments to count result, so the notation is violated
	var expectedErr = "Невірна форма запису виразу! Має використовуватись постфіксна форма!"
	var out bytes.Buffer
	handler := &ComputeHandler{
		Input: strings.NewReader(input),
		Output: &out,
	}
	err := handler.Compute()
	c.Assert(err, gocheck.ErrorMatches, expectedErr)
}

func (s *MySuite) TestComputeWrongOperator_1(c *gocheck.C) {
	var input = "4 5 #" // there is no operator a, so the expression could not be counted
	var expectedErr = "Виконати дану операцію над аргументами неможливо: a!"
	var out bytes.Buffer
	handler := &ComputeHandler{
		Input: strings.NewReader(input),
		Output: &out,
	}
	err := handler.Compute()
	c.Assert(err, gocheck.ErrorMatches, expectedErr)
}

func ExampleCompute() {
	var input = "2 4 ^"
	var out bytes.Buffer
	handler := &ComputeHandler{
		Input: strings.NewReader(input),
		Output: &out,
	}
	err := handler.Compute()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(strings.Trim(out.String(), "\n"))
	}

	// Output:
	// 16
}
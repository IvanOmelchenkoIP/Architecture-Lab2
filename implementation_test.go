package lab2

import (
	"fmt"
	"testing"

	gocheck "gopkg.in/check.v1"
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type MySuite struct{}

var _ = gocheck.Suite(&MySuite{})

func (s *MySuite) TestCountPostfixOperandComplexity1_1(c *gocheck.C) {
	var input = "2" // 2 = 2
	var expected = "2"
	res, err := CountPostfix(input)
	if err != nil {
		c.Fail()
	}
	c.Assert(res, gocheck.Equals, expected)
}

func (s *MySuite) TestCountPostfixOperandComplexity2_1(c *gocheck.C) {
	var input = "3 4 *" // 3 * 4 = 12
	var expected = "12"
	res, err := CountPostfix(input)
	if err != nil {
		c.Fail()
	}
	c.Assert(res, gocheck.Equals, expected)
}

func (s *MySuite) TestCountPostfixOperandComplexity2_2(c *gocheck.C) {
	var input = "12 6 /" // 12 / 6 = 2
	var expected = "2"
	res, err := CountPostfix(input)
	if err != nil {
		c.Fail()
	}
	c.Assert(res, gocheck.Equals, expected)
}

func (s *MySuite) TestCountPostfixOperandComplexity3_1(c *gocheck.C) {
	var input = "18 5 2 ^ +" // 2^5 + 18 = 50
	var expected = "43"
	res, err := CountPostfix(input)
	if err != nil {
		c.Fail()
	}
	c.Assert(res, gocheck.Equals, expected)
}

func (s *MySuite) TestCountPostfixOperandComplexity3_2(c *gocheck.C) {
	var input = "12 8 2 * +" // 12 + 8 * 2 = 28
	var expected = "28"
	res, err := CountPostfix(input)
	if err != nil {
		c.Fail()
	}
	c.Assert(res, gocheck.Equals, expected)
}

func (s *MySuite) TestCountPostfixOperandComplexity4_1(c *gocheck.C) {
	var input = "4 2 - 3 * 5 +" // 5 + (4-2) * 3 = 11
	var expected = "11"
	res, err := CountPostfix(input)
	if err != nil {
		c.Fail()
	}
	c.Assert(res, gocheck.Equals, expected)
}

func (s *MySuite) TestCountPostfixOperandComplexity5_1(c *gocheck.C) {
	var input = "2 5 2 ^ + 9 / 4 +" // (5^2 + 2) / 9 + 4 = 7
	var expected = "7"
	res, err := CountPostfix(input)
	if err != nil {
		c.Fail()
	}
	c.Assert(res, gocheck.Equals, expected)
}

func (s *MySuite) TestCountPostfixOperandComplexity7_1(c *gocheck.C) {
	var input = "4 1 + 2 4 6 + * 6 3 + - *" //(4 + 1) * ((2 * (4 + 6)) - (6 + 3)) = 55
	var expected = "55"
	res, err := CountPostfix(input)
	if err != nil {
		c.Fail()
	}
	c.Assert(res, gocheck.Equals, expected)
}

func (s *MySuite) TestCountPostfixOperandComplexity7_2(c *gocheck.C) {
	var input = "7 3 + 5 3 - / 16 4 + -10 / *" //(7 + 3) / (5 - 3) * ((16 + 4) / (-10)) = -10
	var expected = "-10"
	res, err := CountPostfix(input)
	if err != nil {
		c.Fail()
	}
	c.Assert(res, gocheck.Equals, expected)
}

func (s *MySuite) TestCountPostfixOperandComplexity10_1(c *gocheck.C) {
	var input = "3 2 ^ 4 * 4 - -20 2 * 4 2 ^ + -4 -2 * - /" // (3^2 * 4 - 4) / (-20 * 2 + 4^2 - (-2) * (-4)) = -1
	var expected = "-1"
	res, err := CountPostfix(input)
	if err != nil {
		c.Fail()
	}
	c.Assert(res, gocheck.Equals, expected)
}

func (s * MySuite) TestCountPostfixWrongInput_1(c *gocheck.C) {
	var input = "a" // there is only 2 argument and it is not a number, so it can not be result
	var expectedRes = ""
	var expectedErr = "Невірна форма запису виразу! Має використовуватись постфіксна форма!"
	res, err := CountPostfix(input)
	c.Assert(res, gocheck.Equals, expectedRes)
	c.Assert(err, gocheck.ErrorMatches, expectedErr)
}

func (s * MySuite) TestCountPostfixWrongInput_2(c *gocheck.C) {
	var input = "3 +" // there is not enough arguments to count, so it counts as a postfix notation violation
	var expectedRes = ""
	var expectedErr = "Невірна форма запису виразу! Має використовуватись постфіксна форма!"
	res, err := CountPostfix(input)
	c.Assert(res, gocheck.Equals, expectedRes)
	c.Assert(err, gocheck.ErrorMatches, expectedErr)
}

func (s * MySuite) TestCountPostfixWrongInput_3(c *gocheck.C) {
	var input = "4 3 2 +" // there are not enough operators to process all the operands, so the notation is violated
	var expectedRes = ""
	var expectedErr = "Невірна форма запису виразу! Має використовуватись постфіксна форма!"
	res, err := CountPostfix(input)
	c.Assert(res, gocheck.Equals, expectedRes)
	c.Assert(err, gocheck.ErrorMatches, expectedErr)
}

func (s * MySuite) TestCountPostfixWrongInput_4(c *gocheck.C) {
	var input = "7 a" // though the operator is wrong, there are not enough arguments to count result, so the notation is violated
	var expectedRes = ""
	var expectedErr = "Невірна форма запису виразу! Має використовуватись постфіксна форма!"
	res, err := CountPostfix(input)
	c.Assert(res, gocheck.Equals, expectedRes)
	c.Assert(err, gocheck.ErrorMatches, expectedErr)
}

func (s * MySuite) TestCountPostfixWrongOperator_1(c *gocheck.C) {
	var input = "4 3 a" // there is no operator a, so the expression could not be counted
	var expectedRes = ""
	var expectedErr = "Виконати дану операцію над аргументами неможливо: a!"
	res, err := CountPostfix(input)
	c.Assert(res, gocheck.Equals, expectedRes)
	c.Assert(err, gocheck.ErrorMatches, expectedErr)
}

func ExampleCountPostfix() {
	var input = "2 2 + 4 2 * 3 - +"
	res, err := CountPostfix(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	// Output:
	// 9
}

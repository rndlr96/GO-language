// Eval returns the evaluation result of the givien expr
// The expression can have +, -, *, /, (, ) operators and
// decimal integers. Operators and operands should be space
// delimited

package Eval

import (
	"strconv"
	"strings"
)

func Eval(expr string) int {
	var ops []string
	var nums []int

	// pop is function literal
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	reduce := func(higher string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			// check op exist in expr
			if strings.Index(higher, op) < 0 {
				return
			}
			ops = ops[:len(ops)-1] // remove (
			if op == "(" {
				return
			}
			b, a := pop(), pop()
			switch op {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a-b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)
			}
		}
	}
	for _, token := range strings.Split(expr, " ") {
		switch token {
		case "(":
			ops = append(ops, token)
		case "+", "-":
			reduce("+-*/")
			ops = append(ops, token)
		case "*", "/":
			reduce("*/")
			ops = append(ops, token)
		case ")":
			reduce("+-*/(")
		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	reduce("+-*/")
	return nums[0]
}

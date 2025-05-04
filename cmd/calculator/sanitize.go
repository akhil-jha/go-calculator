package main

import (
	"strconv"
	"strings"
	"unicode"
)

func isDigit(s string) bool {
	for _, i := range s {
		if unicode.IsDigit(i) {
			return true
		}
	}
	return false
}

func SanitizeArg(args []string) (string, int, int) {
	var operator string
	var operand1 int
	var operand2 int
	joined := strings.Join(args, "")
	// Get the operator
	for _, i := range joined {
		if !isDigit(string(i)) {
			operator = string(i)
			break
		}
	}

	//split the string by operator
	operands := strings.Split(joined, operator)
	operand1, _ = strconv.Atoi(operands[0])
	operand2, _ = strconv.Atoi(operands[1])
	return operator, operand1, operand2
}

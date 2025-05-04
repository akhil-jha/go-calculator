package main

import (
	"fmt"
	"os"
	"strconv"

	"local.example.com/go-calculator/go_usage"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Need more args", args)
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("Too many args", go_usage.Usage())
	} else {
		operator, operand1, operand2 := SanitizeArg(args)
		var result string
		switch operator {
		case "+":
			result = strconv.Itoa(operand1 + operand2)
			fmt.Println("Result: ", result)
		case "-":
			result = strconv.Itoa(operand1 - operand2)
			fmt.Println("Result: ", result)
		case "x":
			result = strconv.Itoa(operand1 * operand2)
			fmt.Println("Result: ", result)
		case "/":
			if operand2 == 0 {
				result = "Undefined"
				fmt.Println("Result:", result)
			} else {
				result = strconv.Itoa(operand1 / operand2)
				fmt.Println("Result:", result)
			}
		}
	}
}

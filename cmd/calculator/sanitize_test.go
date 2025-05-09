package main

import (
	"testing"
)

func TestIsDigit(test *testing.T) {
	input := "1"
	output := true

	result := isDigit(input)

	if result != output {
		test.Errorf("Function not working. Expected %v, got %v as input was %v", output, result, input)
	}
}

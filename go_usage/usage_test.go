package go_usage

import (
	"testing"
)

func TestUsage(test *testing.T) {
	expected := "\n\nUsage: go run ./cmd/calculator operand1operatoroperand2"

	if Usage() != expected {
		test.Error()
	}
}

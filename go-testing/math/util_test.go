package math_test

import (
	"go-testing/math"
	"testing"
)

func TestDoMath(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
		errMsg   string
	}{
		{"addition", 2, 2, "+", 4, ""},
		{"subtraction", 2, 2, "-", 0, ""},
		{"multiplication", 2, 2, "*", 4, ""},
		{"division", 2, 2, "/", 1, ""},
		{"bad_division", 2, 0, "/", 0, `division by zero`},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := math.DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error message `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}

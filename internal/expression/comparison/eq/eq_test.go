package eq

import (
	"fmt"
	. "goillogical/internal"
	. "goillogical/internal/mock"
	"testing"
)

func e(val any) Evaluable {
	return E(val, fmt.Sprintf("%v", val))
}

func TestHandler(t *testing.T) {
	var tests = []struct {
		left     Evaluable
		right    Evaluable
		expected bool
	}{
		// Same types
		{e(1), e(1), true},
		{e(1.1), e(1.1), true},
		{e("1"), e("1"), true},
		{e(true), e(true), true},
		{e(false), e(false), true},
		// Diff types
		{e(1), e(1.1), false},
		{e(1), e("1"), false},
		{e(1), e(true), false},
		{e(1.1), e("1"), false},
		{e(1.1), e(true), false},
		{e("1"), e(true), false},
		// Slices
		{e([]int{1}), e([]int{1}), false},
		{e(1), e([]int{1}), false},
	}

	for _, test := range tests {
		c, _ := New("==", test.left, test.right)
		if output, err := c.Evaluate(map[string]any{}); output != test.expected || err != nil {
			t.Errorf("input (%v, %v): expected %v, got %v/%v", test.left.String(), test.right.String(), test.expected, output, err)
		}
	}
}

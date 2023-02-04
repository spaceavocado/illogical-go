package not

import (
	"errors"
	"fmt"
	. "goillogical/internal"
	"testing"
)

type mock struct {
	val any
}

func (m mock) String() string {
	return fmt.Sprintf("%v", m.val)
}

func (m mock) Evaluate(ctx Context) (any, error) {
	return m.val, nil
}

func e(val any) Evaluable {
	return mock{val}
}

func TestHandler(t *testing.T) {
	var tests = []struct {
		operand  Evaluable
		expected bool
	}{
		{e(true), false},
		{e(false), true},
	}

	for _, test := range tests {
		c, _ := New(test.operand)
		if output, err := c.Evaluate(map[string]any{}); output != test.expected || err != nil {
			t.Errorf("input (%v): expected %v, got %v/%v", test.operand, test.expected, output, err)
		}
	}

	var errs = []struct {
		operand  Evaluable
		expected error
	}{
		{e("bogus"), errors.New("logical NOT expression's operand must be evaluated to boolean value")},
	}

	for _, test := range errs {
		c, _ := New(test.operand)
		if _, err := c.Evaluate(map[string]any{}); err.Error() != test.expected.Error() {
			t.Errorf("input (%v): expected %v, got %v", test.operand, test.expected, err)
		}
	}
}

package present

import (
	. "goillogical/internal"
	c "goillogical/internal/expression/comparison"
)

func handler(evaluated []any) bool {
	return evaluated[0] != nil
}

func New(e Evaluable) (Evaluable, error) {
	return c.New("<is present>", []Evaluable{e}, handler)
}

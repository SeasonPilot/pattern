package strategy_pattern

import (
	"fmt"
	"testing"
)

func TestSetStrategy(t *testing.T) {
	o := Operator{}

	o.SetStrategy(&Add{})
	fmt.Printf("%#v\n", o.Calculate(2, 2))

	o.SetStrategy(&Reduce{})
	fmt.Printf("%#v\n", o.Calculate(2, 2))
}

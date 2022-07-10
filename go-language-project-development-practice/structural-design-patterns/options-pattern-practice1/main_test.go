package options_pattern_practice1

import (
	"fmt"
	"testing"
)

func TestNewConn(t *testing.T) {
	cache := WithCache(true)
	fmt.Printf("%#v\n", NewConn("18471", cache))
}

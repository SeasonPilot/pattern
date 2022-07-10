package options_pattern_practice2

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	cache := WithCache(true)
	conn, err := NewConn("980", cache)
	if err != nil {
		return
	}
	fmt.Println(conn)
}

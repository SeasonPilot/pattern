package main

import (
	"fmt"
	"testing"
)

func TestCacheFactory_Create(t *testing.T) {
	cf := CacheFactory{}
	r, err := cf.Create(redis)
	if err != nil {
		t.Error(err)
	}

	r.Set("k1", "v1")
	fmt.Println(r.Get("k1"))
}

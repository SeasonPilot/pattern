package singleton

import "sync"

type singleton struct {
}

var (
	ins  *singleton
	once sync.Once
)

// GetIns The exported function has an unexported return type
func GetIns() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}

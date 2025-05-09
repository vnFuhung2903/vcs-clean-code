package singleton

import (
	"sync"
)

var (
	once    sync.Once
	counter = 0
)

func Count() int {
	once.Do(func() {
		counter++
	})
	return counter
}

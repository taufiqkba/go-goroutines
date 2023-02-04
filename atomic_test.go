package gogoroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// Case Race Condition
func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		group.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				// x += 1 // replace to atomic like bellow
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("Data X - ", x)
}

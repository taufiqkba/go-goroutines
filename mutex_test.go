package gogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex // Initialize to using Mutex (Mutual Exclusion)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock() //Lock first
				x += 1
				mutex.Unlock() // then unlock
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Data X -", x)
}

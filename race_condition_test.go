package gogoroutines

import (
	"fmt"
	"testing"
	"time"
)

// Case Race Condition
func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Data X - ", x)
}

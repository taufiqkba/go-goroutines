package gogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		// Make default value if pool is empty to override function
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Taufiq")
	pool.Put("Kurniawan")
	pool.Put("Bayu")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Done!")
}

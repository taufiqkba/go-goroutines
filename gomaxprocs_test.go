package gogoroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	// Test to create goroutine to make sure number goroutines true
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	cpuTotal := runtime.NumCPU()
	fmt.Println("CPU Total: ", cpuTotal)

	runtime.GOMAXPROCS(20)
	threadTotal := runtime.GOMAXPROCS(-1)
	fmt.Println("THREAD Total: ", threadTotal)

	goroutineTotal := runtime.NumGoroutine()
	fmt.Println("Goroutine Total: ", goroutineTotal)
	group.Wait()
}

// Test Change Thread Numbers
func TestChangeThreadNumber(t *testing.T) {
	// Test to create nil goroutine to make sure num goroutine true
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	cpuTotal := runtime.NumCPU()
	fmt.Println("CPU Total: ", cpuTotal)

	runtime.GOMAXPROCS(20) // To change Thread Number in golang
	threadTotal := runtime.GOMAXPROCS(-1)
	fmt.Println("THREAD Total: ", threadTotal)

	goroutineTotal := runtime.NumGoroutine()
	fmt.Println("Goroutine Total: ", goroutineTotal)
	group.Wait()
}

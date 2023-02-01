package gogoroutines

import (
	"fmt"
	"testing"
	"time"
)

// Create function using golang goroutine
func RunHelloWorld() {
	fmt.Println("Hello World!")
}
func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // using method "go" in front of the function that was wants to implemetn goroutine
	fmt.Println("Uups")

	time.Sleep(1 * time.Second)
}

// Test many goroutine
func DisplayNumber(number int) {
	fmt.Println("Number-", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}

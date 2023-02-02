package gogoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// Create Channel

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Taufiq Kurniawan"
		fmt.Println("Finished send data to Channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// Channel as a parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Give me a Response"
}
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// Channel In and Out
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Channel only IN"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Buffered Channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Taufiq"
		channel <- "Kurniawan"
		channel <- "Baayu"

	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)

	}()
	time.Sleep(2 * time.Second)
	fmt.Println("DONE!")
}

// Range Channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Looping for - " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Receive data,", data)
	}
}

// Select Channel

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from Channel 1,", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from Channel 2,", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}

}

// Default Channel
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from Channel 1,", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from Channel 2,", data)
			counter++
		default:
			fmt.Println("Waiting data response...")
		}

		if counter == 2 {
			break
		}
	}

}

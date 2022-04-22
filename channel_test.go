package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "Hello, World!"
	}()

	data := <-ch
	fmt.Println(data)

	defer close(ch)
}

func TestChannelAsParameter(t *testing.T) {
	ch := make(chan string)

	go GiveMeResponse(ch)

	data := <-ch
	fmt.Println(data)

	defer close(ch)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(1 * time.Second)

	channel <- "Hello, World!"
}

func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "Hello, World!"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	ch := make(chan string)

	go OnlyIn(ch)
	go OnlyOut(ch)

	time.Sleep(3 * time.Second)
	defer close(ch)
}

func TestCreateBufferedChannel(t *testing.T) {
	ch := make(chan string, 3)

	defer close(ch)

	ch <- "Hello, World!"
	ch <- "Hello, World!"
	ch <- "Hello, World!"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("Done")
}

func TestRangeChannel(t *testing.T) {
	ch := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("Hello, World! %d", i)
		}
		defer close(ch)
	}()

	for data := range ch {
		fmt.Println(data)
	}

	fmt.Println("Done")
}

func TestSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	defer close(ch1)
	defer close(ch2)

	go GiveMeResponse(ch2)
	time.Sleep(1 * time.Second)
	go GiveMeResponse(ch1)

	counter := 0
	for {
		select {
		case data := <-ch1:
			fmt.Println(data, "1")
			counter++
		case data := <-ch2:
			fmt.Println(data, "2")
			counter++
		default:
			fmt.Println("Nothing")
		}

		if counter == 2 {
			break
		}
	}
}

package main

import (
	"fmt"
)

func send(ch chan<- string) {
	for i := 1; i <= 20; i++ {
		temp := fmt.Sprintf("Hello i ke-%d", i)
		ch <- temp
	}
	close(ch)
}

func receipe(ch <-chan string) {
	for each := range ch {
		fmt.Println(each)
	}
}

func main() {
	ch := make(chan string)
	go send(ch)
	receipe(ch)

	chh := make(chan string)
	go func() {
		chh <- "Test test"
	}()
	fmt.Println(<-chh)
}

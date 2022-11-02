package main

import "fmt"

func main() {
	chanel := make(chan string)

	clos := func(who string) {
		temp := fmt.Sprintf("Hello %s", who)
		chanel <- temp
	}

	go clos("Muhammad")
	go clos("Nawi")
	go clos("Husen")
	depan := <-chanel
	fmt.Println(depan)

	tengah := <-chanel
	fmt.Println(tengah)

	last := <-chanel
	fmt.Println(last)
}

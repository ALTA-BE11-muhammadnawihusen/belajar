package main

import (
	"fmt"
)

func keluar(plus <-chan string) {
	result := "Aku adalah Pencari Receh"
	fmt.Println(<-plus, result)
	fmt.Println("Apakah ini masuk")
	// test := <-plus //akan error karena tidak ada lagi isi di dalam channel
	// fmt.Println(test)
}

func masuk(plus chan<- string, say string) {
	plus <- say
}

func main() {

	// runtime.GOMAXPROCS(1)
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

	fmt.Println("======================")

	go masuk(chanel, "Hello I am Nawi Husen")
	go masuk(chanel, "Hello I am Jack")
	// test2 := <-chanel
	// fmt.Println(<-chanel)
	// fmt.Println(<-chanel)
	// fmt.Println(test2)
	// fmt.Println(test2)
	keluar(chanel)
	keluar(chanel)
}

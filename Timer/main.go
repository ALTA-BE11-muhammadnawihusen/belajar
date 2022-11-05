package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Sleep(5 * time.Second)
	// fmt.Println("Ini di print setelah 5 detik")

	// chtime := time.NewTimer(5 * time.Second)
	// fmt.Println("Hello")

	// <-chtime.C

	// fmt.Println("Sampai Jumpa")
	fmt.Println("=========================")
	ch := make(chan string)

	say := func() {
		ch <- "Hello"
	}

	time.AfterFunc(5*time.Second, say)

	time.Sleep(1 * time.Second)
	fmt.Println("Test 1")
	time.Sleep(1 * time.Second)
	fmt.Println("Test 2")
	//di sini karena akan mengunngu ch dia akan diam lebih dari satu detik tergantung seberapa jauh dari titik mulainya
	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)
	fmt.Println("Test 3")
	time.Sleep(1 * time.Second)
	fmt.Println("Test 4")
	time.Sleep(1 * time.Second)
	fmt.Println("Test 5")
	time.Sleep(1 * time.Second)
	fmt.Println("Test 6")

	fmt.Println("=====================")

	<-time.After(5 * time.Second)
	fmt.Println("Test 7")
}

// hampir sama seperti time.sleep tapi type data nya adalah channel
// time.Afterfunc bersifat assyncronous jadi tidak akan menunggu dan bisa saja tidak akan di jalankan kecuali ada channel
// maka proses akan terhenti sementara pada penerimaan channel
// time.after sama persis seperti sleep tapi penerapannya menggunakan channel

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func catch() {
	s := recover()
	if s != nil {
		fmt.Println("Terdapat sebuah error di dalam code")
	} else {
		fmt.Println("Aplikasi berjalan lancar")
	}
}

func main() {

	var input string
	// var some error
	fmt.Scan(&input)
	defer catch()

	num, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println(err)
		// fmt.Println(some.Error())
		panic("tolong masukkan nomor yang benar")
	} else if err == nil {
		fmt.Println(num)
	}

	tes := errors.New("Ini adalah error yang di buat")
	fmt.Println(tes)

}

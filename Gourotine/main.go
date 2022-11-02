package main

import (
	"fmt"
	"runtime"
)

func print(jumlah int, kata string) {
	for i := 1; i <= jumlah; i++ {
		fmt.Println(kata, i)
	}
}

func pri(s string, ke int) {
	fmt.Println(s, ke)
}

func main() {
	runtime.GOMAXPROCS(2)
	// for i := 0; i <= 10000; i++ {
	// 	go pri("coba ke-", i)
	// } akan menghasilkan angka yang tidak terurut karena go pri tidak akan menunggu
	// yang seharusnya berjalan lebih dlu

	go print(1000, "percobaan ke")
	print(1000, "kemampuan ke")

	// var test string
	// fmt.Scan(&test)
}

//	runtime.GOMAXPROCS(n) digunakan untuk menentukan jumlah core yang diaktifkan saat eksekusi
// program

package main

import "fmt"

func main() {
	// biasa
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// hanya kondisi
	var i = 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	// tanpa argument
	i = 0
	for {
		fmt.Println("Angka", i)
		i++
		if i == 5 {
			break
		}
	}
}

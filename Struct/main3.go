package main

import "fmt"

type Person struct {
	Nama string
	Age  int
}

func main() {
	orang := &Person{
		Nama: "Nawi",
		Age:  22,
	}

	orang1 := orang
	orang.Nama = "Jack"
	orang1.Nama = "Reaper"

	fmt.Println(orang)
}

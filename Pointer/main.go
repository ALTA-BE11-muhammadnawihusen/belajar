package main

import "fmt"

type Person struct {
	Nama, Address string
	Age           uint
}

func Get(nama, address string, age uint) *Person {
	return &Person{
		Nama:    nama,
		Address: address,
		Age:     uint(age),
	}
}

func Got(name *string) string {
	return *name
}

func main() {
	nama := "Muhammad Nawi Husen"
	fmt.Println("Nama sebelem Berubah", nama)
	alamatnama := &nama
	*alamatnama = "Jack Reaper"

	fmt.Println(nama)
	fmt.Println(&nama)
	fmt.Println(alamatnama)
	fmt.Println(*alamatnama)
	nama = "Nawi"
	fmt.Println(nama)
	fmt.Println(*alamatnama)
	fmt.Println("===========================")

	orang1 := Get("Nawi", "Japan", 22)
	fmt.Println(orang1)
	fmt.Println(*orang1)
	fmt.Println(&orang1)

	namabaru := "Jack"

	namae := Got(&namabaru)
	fmt.Println(namae)

}

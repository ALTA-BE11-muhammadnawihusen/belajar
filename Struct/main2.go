package main

import "fmt"

type Person struct {
	Nama, Address string
}

type Student struct {
	Grade  int
	person Person
}

func main() {
	orang := Person{
		Nama:    "Nawi",
		Address: "Jepang",
	}

	murid := Student{
		Grade:  6,
		person: orang,
	}

	fmt.Println(murid)
	murid.person.Nama = "Jack"
	fmt.Println(murid)

	//kita juga bisa menggunakan alias
	type Hito = Person

	hito1 := Hito{
		Nama:    "Sakura",
		Address: "Kyoto",
	}

	fmt.Println(hito1)
}

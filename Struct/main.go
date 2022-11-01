package main

import "fmt"

type Person struct {
	Name, Address string
	Age           uint
	Hobbies       []string
}

func main() {
	orang1 := Person{
		Name:    "Nawi",
		Address: "Jepang",
		Age:     22,
		Hobbies: []string{"Playing Game", "Read Novel"},
	}

	// var orang2 Person
	// orang2.Name = "Husen"
	// orang2.Address = "Russia"
	// orang2.Age = 30
	// orang2.Hobbies = append(orang2.Hobbies, "All Of Thing")

	fmt.Println(orang1)
	// fmt.Println(orang2)

	orang3 := &orang1
	// meskipun orang3 adalah address. jika ingin mengakses nilai yang ada bisa langsung saja
	fmt.Println("nama orang 1", orang3.Name)
	orang3.Name = "Jack"

	fmt.Println(orang3)
	fmt.Println(orang1)
}

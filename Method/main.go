package main

import "fmt"

type Person struct {
	Nama, Address string
	Age           int
}

func (P Person) Sapa(To string) {
	fmt.Printf("Hello %s My Name is %s,Nice to meet you \n", To, P.Nama)
}

func (P Person) SapaR(To string) string {
	return fmt.Sprintf("Hello %s I live in %s", To, P.Address)
}

func main() {
	Nawi := Person{
		Nama:    "Nawi",
		Address: "Jepang",
		Age:     22,
	}

	Nawi.Sapa("Agus")
	alamatnawi := Nawi.SapaR("Agus")
	fmt.Println(alamatnawi)
}

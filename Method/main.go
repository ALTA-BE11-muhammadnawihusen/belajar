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

//meskipun terjadi perubahan di block fungsi nilai pada properti itu sendiri masih tetap
func (P Person) MakeChange(new string) {
	P.Nama = new
	fmt.Println("Ganti Jadi", P.Nama)
}

// dengan menggunakan pointer pada method kita akan dapat memanipulasi property yang di miliki nya
func (P *Person) MakeChangeP(new string) {
	P.Nama = new
	fmt.Println("Ganti Jadi", P.Nama)
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

	fmt.Println("=========================")

	Nawi.MakeChange("Jack")
	fmt.Println(Nawi.Nama)

	fmt.Println(Nawi)
	fmt.Println("========================")

	Nawi.MakeChangeP("Jack")
	fmt.Println(Nawi.Nama)

	fmt.Println(Nawi)

}

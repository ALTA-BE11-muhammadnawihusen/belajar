package paket

import f "fmt"

type Person struct {
	Nama string
	Age  int
}

var Orang Person

// kita dapat membuat banyak sekali func init sebanyak yang kita mau

func init() {
	Orang.Nama = "Nawi"
	Orang.Age = 22

	f.Println("Ini adalah fungsi milik init")
	f.Println("Ini adalah init 1")
}

func init() {
	f.Println("Saya ingin belajar banyak hal")
	f.Println("Karena itu lah saya berusaha keras")
	f.Println("dengan semua yang saya lakukan")
	f.Println("Ini adalah init 2")
}

package main

import "fmt"

func main() {
	// any == interface{}
	var baru any
	baru = 15
	fmt.Println(baru)
	baru = "Ini adalah nilai baru"
	fmt.Println(baru)
	baru = []string{"berubah", "menjadi", "array"}
	fmt.Println(baru)
	fmt.Println("===================")
	var an any
	an = 10
	result := an.(int) * 10
	fmt.Println(result)
	an = "String"
	result2 := an.(string) + " Terbaru"
	fmt.Println(result2)
	an = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result3 := len(an.([]int))
	fmt.Println(result3)
}

// PENTING
// Apa yang dimiliki interface{}/any hanyalah semacam tiruan yang ditampilkan
// untuk menjadikannya asli perlu di lakukan casting sesuai tipe yang di muatnya

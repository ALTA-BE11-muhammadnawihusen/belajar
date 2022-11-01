package main

import "fmt"

type Count interface {
	Luas() float64
	Keliling() float64
}

type Square struct {
	Panjang float64
}

type lingkaran struct {
	Diameter float64
}

func (s Square) Luas() float64 {
	return float64(s.Panjang * s.Panjang)
}

func (s Square) Keliling() float64 {
	return s.Panjang * 4
}

func main() {
	// kotak := Square{Panjang: 6}
	// luaskotak := kotak.Luas()
	// kelilingkotak := kotak.Keliling()
	// fmt.Println(luaskotak)
	// fmt.Println(kelilingkotak)
	var test Count
	test = Square{Panjang: 5}
	fmt.Println(test.Keliling())
	fmt.Println(test.Luas())
}

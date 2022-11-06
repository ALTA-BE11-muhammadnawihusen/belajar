package main

import (
	"flag"
	"fmt"
)

func main() {
	add := flag.String("Address", "Japan", "Write your address")
	// nilai kembalian dari flag adalah pointer jadi jika ingin mengaksesnya
	// harus meggunakan *
	// parameter pertama berisi key, yang kedua berisi default value, yang ketiga berisi semacam pesan
	age := flag.Int64("Age", 18, "Write your Age")

	// kita juga bisa menggunakan cara seperti dobawah tapi lebih sulit
	// tambahkan akhiran Var setelah type datanya
	var data2 string
	flag.StringVar(&data2, "gender", "male", "type your gender")
	flag.Parse()

	fmt.Println(*add)
	fmt.Println(*age)
	fmt.Println(data2)

}

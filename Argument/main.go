package main

import (
	"fmt"
	"os"
)

func main() {
	test := os.Args
	fmt.Println(test)

	//os.Args menghasilkan []string dimulai dari index ke 0
	// argument di masukkan saat menjalankan program,misal go run main.go arg1 arg2 arg3
	// argument akan dipisahkan dengan spasi
	// jika sebuah argument berisi spasi maka harus diapit " agar dihitug satu kesatuan
	test2 := os.Args[2]
	fmt.Println(test2)

	test3 := os.Args[3]
	fmt.Println(test3)

	test4 := os.Args[4]
	fmt.Println(test4)

	arg := test[1:]
	fmt.Println(arg)

}

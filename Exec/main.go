package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Println("Exec")
	test, _ := exec.Command("ls").Output() // menghasilkan []byte
	fmt.Println(string(test))

	fmt.Println("==================")
	test2, _ := exec.Command("pwd").Output() // menghasilkan []byte
	fmt.Println(string(test2))

	fmt.Println("==================")
	test3, _ := exec.Command("git", "config", "user.name").Output() // bisa gunakan koma untuk pemisah
	// test4, _ := exec.Command("git config user.name").Output()       //tidak akan menghasilkan apa apa
	fmt.Println(string(test3))
	// fmt.Println(string(test4)) // tidak akan menghasilkan apa apa

	t := runtime.GOOS // mengembalikan sistem operasi yang digunakan dalam bentuk string
	fmt.Println(t)

	// tes, _ := exec.Command("mkdir", "File").Output()
	// fmt.Println(tes) akan membut file
}

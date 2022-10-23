package main

import "fmt"

func main() {

	fmt.Println("john wick")
	fmt.Println("john" + "wick")

	fmt.Print("john wick\n")
	fmt.Print("john", "wick\n")
	fmt.Print("john", " ", "wick\n")

	const name string = "Nawi"
	const age = 22
	const stat = true
	const tes1, tes2, tes3 = 1, "tes", 3
	fmt.Println(tes1, tes2, tes3)

}

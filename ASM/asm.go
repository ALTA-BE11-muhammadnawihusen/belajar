package main

import "fmt"

func main() {
	//salah satu cara membuat array
	var fruits = make([]string, 2)
	fruits[0] = "apple"
	fruits[1] = "manggo"

	fmt.Println(fruits) // [apple manggo]
}

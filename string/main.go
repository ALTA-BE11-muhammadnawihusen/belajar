package main

import (
	"fmt"
	"strings"
)

func main() {
	// test := strings.HasPrefix("Muhammad Nawi Husen", "Na")
	// fmt.Println(test)

	// test := strings.Count("Muhammad Nawi Husen", "M")
	// fmt.Println(test)

	// test := strings.Index("Muhammad Nawi Husen", "usen")
	// fmt.Println(test)

	// test := strings.Replace("Muhammad Nawi Husen", "a", "i", 1)
	// test := strings.Replace("Muhammad Nawi Husen", "a", "i", -1) // mengganti semua a menjadi i
	// fmt.Println(test)

	// test := strings.Repeat("Allah\n", 10) // mengganti semua a menjadi i
	// fmt.Println(test)

	// test := strings.Split("Muhammad Nawi Husen", "a")
	// fmt.Println(test)

	test := []string{"M", "u", "h", "a", "m", "m", "a", "d"}
	test2 := strings.Join(test, " ")
	fmt.Println(test2)
}

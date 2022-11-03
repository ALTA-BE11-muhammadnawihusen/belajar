package main

import "fmt"

func main() {
	arr := []string{"Muhammad", "Nawi", "Husen"}

	for _, each := range arr {
		func() {

			defer func() {
				if s := recover(); s != nil {
					fmt.Println("Terdapat error")
				} else {
					fmt.Println("everything is alright")
				}
			}()

			panic("ini adalah panic")

		}()
		fmt.Println(each)
	}
}

// ketika panic tidak di recover maka akan menghentikan program secara keseluruhan
// tergantung kebutuhan akan lebih baik untuk memakai recover

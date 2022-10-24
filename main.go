package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	ar := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	arr2 := []int{}

	arr2 = append(arr2, arr...)
	arr2 = append(arr2, ar...)

	fmt.Println(arr)
	fmt.Println(arr2)
}

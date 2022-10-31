package main

import "fmt"

func main() {
	clousure := func(arr []int) (int, int) {
		var max, min int
		for k, v := range arr {
			switch {
			case k == 0:
				min, max = v, v
			case max < v:
				max = v
			case min > v:
				min = v
			}
		}

		return min, max
	}

	tes := []int{1, 4, 6, 89, 6, 3, 463, 3, 2, 11, 0}
	min, max := clousure(tes)
	fmt.Println("ini adalah nilai min", min)
	fmt.Println("ini adalah nilai max", max)
}

package main

import "fmt"

func main() {
	od := 2
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

	tes := []int{1, 4, 6, 89, 232, 234, 44, 14, 45, 12, 6, 3, 463, 3, 2, 11, 0}
	min, max := clousure(tes)
	fmt.Println("ini adalah nilai min", min)
	fmt.Println("ini adalah nilai max", max)

	IIFO1, IIFO2 := func(odd int) ([]int, int) {
		var res []int
		for _, v := range tes {
			if v%odd == 0 {
				res = append(res, v)
			}
		}

		return res, len(res)
	}(od) //kurung di sini bersifat wajib dan akan dijadikan parameter

	fmt.Println("IIFO1 :", IIFO1, "IIFO2 :", IIFO2)

}

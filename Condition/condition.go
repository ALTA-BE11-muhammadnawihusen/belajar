package main

import "fmt"

func main() {

	var point = 10

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	case 9, 10:
		fmt.Println("You are the best")
		fmt.Println("Kamu adalah yang terbaik")
		//{} kamu bisa menambahkan kurung kurawal pada code di atas
	default:
		fmt.Println("not bad")
	}

	switch {
	case point > 1 && point < 5:
		fmt.Println("perfect")
	case point >= 5 && point < 10:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	point = 10

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}

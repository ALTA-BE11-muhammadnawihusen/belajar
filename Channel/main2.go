package main

import (
	"fmt"
)

// func main() {
// 	channel := make(chan string)
// 	tes := 6

// 	go func(count int) {
// 		for i := 1; i <= count; i++ {
// 			temp := fmt.Sprintf("Halo aku perulangan ke- %d\n", i)
// 			temp2 := fmt.Sprintf("Halo aku perulangan ke- %d\n", i+count)
// 			channel <- temp
// 			channel <- temp2
// 		}
// 		close(channel)
// 	}(tes)

// 	for each := range channel {
// 		fmt.Println(each)
// 	}

// 	// fmt.Println(<-channel)
// 	// fmt.Println(<-channel)
// 	// fmt.Println(<-channel)
// 	// fmt.Println(<-channel)
// 	// fmt.Println(<-channel)
// 	// fmt.Println(<-channel)
// }

func main() {
	channel := make(chan string, 2)
	defer close(channel)

	go func(pesan string) {
		for i := 1; i <= 5; i++ {
			fmt.Println("Ini masuk ke channel ke-", i)
			temp := fmt.Sprintf("Ini adalah %s ke- %d", pesan, i)
			channel <- temp
		}
	}("Hello")

	for i := 1; i <= 5; i++ {
		fmt.Println(<-channel)
	}

}

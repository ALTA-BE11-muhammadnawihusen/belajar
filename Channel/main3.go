package main

import (
	"fmt"
	"sort"
	"strconv"
)

func Avg(ch chan<- float64, ar ...int) {
	temp := 0
	for _, value := range ar {
		temp += value
	}
	men := fmt.Sprintf("%.2f", float64(temp)/float64(len(ar)))
	avg, _ := strconv.ParseFloat(men, 64)
	ch <- avg
}

func Min(ch chan<- int, ar ...int) {
	sort.SliceStable(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})

	min := ar[0]
	ch <- min
}

func Max(ch chan<- int, ar ...int) {
	sort.SliceStable(ar, func(i, j int) bool {
		return ar[i] > ar[j]
	})

	max := ar[0]
	ch <- max
}

func main() {
	arr := []int{3, 54, 23, 56, 335, 7, 832, 453, 23, 562, 0, 53, 5752, 1}

	chavg := make(chan float64)
	go Avg(chavg, arr...)

	chmin := make(chan int)
	go Min(chmin, arr...)

	chmax := make(chan int)
	go Max(chmax, arr...)

	for i := 1; i <= 3; i++ {
		select {
		case avg := <-chavg:
			fmt.Println("ini adalah nilai rata rata", avg)
		case min := <-chmin:
			fmt.Println("ini adalah nilai terkecil", min)
		case max := <-chmax:
			fmt.Println("ini adalah nilai terbesar", max)
			// goroutine manapun yang akan selesai lebih dlu maka akan di eksekusi,
			//tidak diketahui mana yang lebih dlu
			// jumlah di perulangan di sesuaikan dengan goroutine yang masuk
		}
	}
	// fmt.Println(<-chavg)
	// fmt.Println(<-chmin)
	// fmt.Println(<-chmax)
}

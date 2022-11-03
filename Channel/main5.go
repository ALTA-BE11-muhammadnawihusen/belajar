package main

import (
	"fmt"
	"time"
)

func Send(channel chan<- int) {
	fitst := 1
	for {
		channel <- fitst
		time.Sleep(time.Duration(fitst) * time.Second)
		fitst += 1
	}
}

func Receipe(channel <-chan int) {
loop:
	for {
		select {
		case data := <-channel:
			fmt.Println(data)
		case <-time.After(time.Second * 5):
			fmt.Println("Kau telah melewati batas waktu")
			break loop
		}
	}
}

func main() {
	chann := make(chan int)

	go Send(chann)
	Receipe(chann)
}

/*

Kondisi case <-time.After(time.Second * 5):, akan terpenuhi ketika tidak ada
aktivitas penerimaan data dari channel dalam durasi 5 detik.
Blok inilah yang kita sebut sebagai blok timeout.
*/

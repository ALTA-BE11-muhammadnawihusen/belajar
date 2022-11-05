package main

import (
	"fmt"
	"time"
)

func main() {
	tes := make(chan bool)
	tic := time.NewTicker(2 * time.Second)

	go func() {
		time.Sleep(20 * time.Second)
		tes <- true
	}()

	for {
		select {
		case <-tes:
			fmt.Println("Program akan dihentikan")
			return
		case <-tic.C:
			fmt.Println("Time Ticker berjalan")
		}
	}
}

//time.ticker dapat digunakan untuk melakukan schejule pada channel untuk memilih yang mana yang akan di jalankan
// yang diatas akan bejalan selama 20 detik dan akan melakukan print time ticker berjalan sebanyak 10 kali
//dan ketika chan tes berjalan maka program akan di hentikan

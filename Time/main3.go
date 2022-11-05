package main

import (
	"fmt"
	"time"
)

// golang juga telah menyediakan beberapa format terlebih dahulu
// meskipun format hasil tetap seperti time.Now()
// lebih baik buat costum sendiri seperti di main2.go

func main() {
	val := fmt.Sprintf("Fri, 14 Aug 2000 10:10:10 WITA")

	// RFC11223 memiliki format Mon, 02 Jan 2006 15:04:05 MST
	date, _ := time.Parse(time.RFC1123, val)
	fmt.Println(date)
}

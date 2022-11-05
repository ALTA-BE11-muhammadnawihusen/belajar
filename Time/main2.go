package main

import (
	"fmt"
	"time"
)

/*
Parsing dari string ke time time. untuk melakukan ini di perlukan untuk membuat format terlebih dahulu sesuai dengan
standard yang ada. untuk hasilnya sendiri pasti akan memiliki format seperti time.Now()
*/

func main() {

	// layoutFormat := "06-Jan-2 03:4:5 PM" jika menggunakan format pm
	// layoutFormat := "06-Jan-2 15:4:5.999999 MST" jika ingin menambahkan lokasi timezone
	layoutFormat := "06-Jan-2 15:4:5.999999 Z0700 MST" // jika ingin menambahkan offsite timezone
	value := fmt.Sprintf("%d-%s-%s %s:%d:%d.%d %s %s", 22, "Aug", "9", "10", time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), "+0800", "WITA")

	fmt.Println(value)
	date, _ := time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date)
	fmt.Println(time.Now())
	fmt.Println("====================")

	form := "2006-January-2" // ketika hanya sebagian maka yang lain akan bernilai 00:00:00
	val := fmt.Sprintf("%v-%v-%v", time.Now().Year(), time.Now().Month(), time.Now().Day())
	fmt.Println(val)
	date2, _ := time.Parse(form, val)

	fmt.Println(date2)

	fmt.Println("==================")
	form = "15:4:5"
	val = fmt.Sprintf("%d:%d:%d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	date3, _ := time.Parse(form, val)
	fmt.Println(date3)
	tim := date3.Minute()
	fmt.Println(tim)

}

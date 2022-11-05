package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Ini adalah penilaian zero ", time.Since(start))
	// menunjukkan waktu yang dibutuhkan dari start

	time.Sleep(1 * time.Second)

	fmt.Println("Ini adalah penilaian pertama ", time.Since(start))

	time.Sleep(1 * time.Second)

	fmt.Println("Ini adalah penilaian kedua ", time.Since(start))

	time.Sleep(1 * time.Second)

	fmt.Println("Ini adalah penilaian ketiga ", time.Since(start))

	time.Sleep(1 * time.Second)

	fmt.Println("Ini adalah penilaian keempat ", time.Since(start))
	fmt.Println("Ini adalah penilaian keempat ", time.Since(start).Seconds())
	fmt.Println("Ini adalah penilaian keempat ", time.Since(start).Minutes())
	fmt.Println("Ini adalah penilaian keempat ", time.Since(start).Hours())

	str := time.Now()

	fmt.Println("Ini adalah selisih waktu antara time.Time1 dan time.Time2", str.Sub(start))
	// sub sendiri akan menghasilkan time.Duration

	//jika kita ingin langsung membuat time.duration itu sendiri mudah, kita bisa membuatnya langsung
	duras := 5 * time.Second //akan menghasilkan time.Duration
	duram := 5 * time.Minute //akan menghasilkan time.Duration
	durah := 5 * time.Hour   //akan menghasilkan time.Duration
	fmt.Println(duras)
	fmt.Println(duram)
	fmt.Println(durah)
	durasm := duras + duram + durah
	fmt.Println("ini adalah gabungan detik menit", durasm.Seconds())
	// hasil penjumlahan dari 5 detik + 5 menit + 5 jam akan dikonversi menjadi detik
}

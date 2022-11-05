package main

import (
	"fmt"
	"time"
)

// melakukan format dari time.Time ke string

func main() {
	lay := "2006-01-02 15:04:05 Z0700 MST"
	val := "2022-11-05 11:08:23 +0800 WITA"
	tes, _ := time.Parse(lay, val)
	// fmt.Println(tes)

	// custom bebas selama sesuai dengan aturan format yang ada
	// tanggal yang di ubah ke sting akan secara otomati mengikuti format yang di buat
	// stdate := tes.Format("Monday 02, January 2006 15:04 MST")
	// stdate := tes.Format("03 02 2006 15:04:05")
	stdate := tes.Format("ini adalah tanggal 03 ini adalah bulan 02 ini adalah tahun 2006 ini adalah jam 15: ini adalah menit 04: ini adalah detik 05")
	fmt.Println(stdate)
}

// untuk penghandle an error bisa langsung di ambil dari hasil kedua dari time.parse atau time.format

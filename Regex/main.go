package main

import (
	"fmt"
	"regexp"
)

func main() {
	regexpat, err := regexp.Compile(`[a-z]+`)
	//maksudnya adalah pola regex akan terdiri dari a sampai z
	text := "Muhammad MaMaM nawi Husen"
	// text := "JUYTRaTYUIJ"
	if err != nil {
		fmt.Println(err.Error())
	}

	test1 := regexpat.FindAllString(text, -1)
	// jumlah hasil dari operasi di atas dpat ditentukan sesuai keiinginan kita
	// hasil akan dipisahkan berdasarkan string yang beruurutan mengikuti sebuah pola yang di tentukan
	// dan akan dipisahkan jika ada string yang keluar dari pola hingga tercipta sebuah slice
	// gunakan -1 untuk melakukan nya sampai akhir
	fmt.Println(test1)

	test2 := regexpat.MatchString(text)
	// selama ada bahkan 1 huruf yang memenuhi pola maka akan bernilai true
	fmt.Println(test2)

	test3 := regexpat.FindString(text)
	fmt.Println(test3)
	//sama seperti yang di atas tapi yang ini hanya bisa mengembalikan satu hasil saja

	test4 := regexpat.FindStringIndex(text)
	fmt.Println(test4)
	// mengembalikan index string dari suatu text yang memenuhi pola tertentu

	test5 := regexpat.ReplaceAllString(text, "Ganti")
	fmt.Println(test5)
	//semua bagian element yang dihasilkan dari find akan dirubah menjadi yang diinginkan sedangkan
	// yang tidak memenuhi pola regex akan tetap tanpa perubahan apapun

	test6 := regexpat.ReplaceAllStringFunc(text, func(each string) string {
		if each == "a" {
			return "pasbanar"
		}
		return each
	})
	fmt.Println(test6)
	// merubah menggunakan fungsi. jika terdapat yang memenuhi kondisi. yang di atas jika hanya == a.
	// jika itu aa maka akan di nilai tidak memenuhi kondisi

	pola, err := regexp.Compile(`[a-b]+`)
	test7 := pola.Split(text, -1)
	fmt.Println(test7)
	//akan mengembalikan slice yang mana pola yang digunakan akan di gunakan sebagai pemisah

	//melakukan percobaan sendiri
	tet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	last, Err := regexp.Compile(`[a-i k-s y-z A-K NOPQZ]+`)
	// tanda - menunjukkan sampai kita bebas membuat pola yang kita inginkan
	if Err != nil {
		fmt.Println(Err.Error())
	}
	last1 := last.FindAllString(tet, -1)

	fmt.Println(last)
	fmt.Println(last1)
}

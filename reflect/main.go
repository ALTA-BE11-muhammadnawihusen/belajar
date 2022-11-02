package main

import (
	"fmt"
	"reflect"
)

func main() {
	test := 33
	reftest := reflect.ValueOf(test)

	fmt.Println(reftest.Type())
	fmt.Println(reftest)

	if reftest.Kind() == reflect.Int {
		temp := reftest.Int()
		fmt.Println("Ini adalah integer :", temp)
	}

	fmt.Println(reflect.Int)
	fmt.Println(reflect.String)
	fmt.Println(reftest.Kind())

	inter := reftest.Interface()
	fmt.Println(inter)
}

/*
object reflect.Value memiliki beberapa method seperti
Kind() untuk mengecek tipe data yang saat ini ditampung sebuah variable
Type() yang mengembalikan type data dari objek yang bersangkutan dalam bentuk string
Int() menghasilkan nilai int dari variable number
String() menghasilkan string
Float64() menghasilkan float64

jika method yang dipanggil tidak sesuai dengan tipe data yang ada didalam reflect.Value
maka akan menyebabkan error

Jika nilai hanya diperlukan untuk ditampilkan ke output, bisa menggunakan
.Interface(). Lewat method tersebut segala jenis nilai bisa diakses dengan mudah.
*/

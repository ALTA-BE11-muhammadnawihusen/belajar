package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Nama string
	Age  int
}

//jika reflect value awal merupakan pointer maka ubahlah agar langsung mengambil elementnya
// untuk memastikan bisa menggunakan reflect.Ptr
// saja dengan menggunakan metod .Elem()

func (s *Person) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)
	fmt.Println("reflectValue :", reflectValue)

	//akan berjalan jika ini merupakan pointer
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	fmt.Println("reflectvalue.elem:", reflectValue)

	//akan menghasilkan type yang sesuai dengan strut yang kita buat
	var reflectType = reflectValue.Type()
	fmt.Println("reflectvalue.type:", reflectType)

	fmt.Println("Number field", reflectValue.NumField())
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name)         // mengembalikan nama property
		fmt.Println("tipe data :", reflectType.Field(i).Type)         // mengembalikan type data property
		fmt.Println("nilai     :", reflectValue.Field(i).Interface()) //menunjukkan isi property
		fmt.Println("")
	}
}

func main() {
	orang := Person{
		Nama: "Nawi",
		Age:  22,
	}

	orang.getPropertyInfo()

}

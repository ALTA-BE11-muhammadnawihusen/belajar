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

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	fmt.Println("reflectvalue.elem:", reflectValue)

	var reflectType = reflectValue.Type()
	fmt.Println("reflectvalue.elem:", reflectType)

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
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

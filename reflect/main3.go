package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Add  string
	Age  uint
}

func (s *Person) Changename(name, Add string, age int) {
	s.Name = name
	s.Add = Add
	s.Age = uint(age)
}

func (s *Person) Changeage(age uint) {
	s.Age = age
}

func main() {
	orang := &Person{Name: "Nawi", Age: 22, Add: "Amuntai"}
	fmt.Println(orang)

	reflectValue := reflect.ValueOf(orang)
	gantinama := reflectValue.MethodByName("Changename")
	gantinama.Call([]reflect.Value{reflect.ValueOf("Jack"), reflect.ValueOf("Jepang"), reflect.ValueOf(int(33))})
	fmt.Println(orang)
}

// ([]reflect.Value{reflect.ValueOf("Baru")}) digunakan untuk mengisi parameter method
// index ke 0 merupakan parameter pertama

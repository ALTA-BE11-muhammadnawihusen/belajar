package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	ar := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	arr2 := []int{}

	arr2 = append(arr2, arr...)
	arr2 = append(arr2, ar...)

	fmt.Println(arr)
	fmt.Println(arr2)
}

// mockery --dir=featur/user(lokasi interface yang akan di mock) --name=Datainterface(nama interface yang di mocking) --filename=UserData.go(nama file hasil generate) --structname=UserData(nama hasil strucr yang di mock)

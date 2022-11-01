package main

import "fmt"

func main() {
	var mm []string
	fmt.Println(mm)
	mm = append(mm, "Baru")
	fmt.Println(mm)

}

// mockery --dir=featur/user(lokasi interface yang akan di mock) --name=Datainterface(nama interface yang di mocking) --filename=UserData.go(nama file hasil generate) --structname=UserData(nama hasil strucr yang di mock)

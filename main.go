package main

import (
	"fmt"
)

func main() {
	str := "az"
	fmt.Println(str[0], str[1])
}

// mockery --dir=featur/user(lokasi interface yang akan di mock) --name=Datainterface(nama interface yang di mocking) --filename=UserData.go(nama file hasil generate) --structname=UserData(nama hasil strucr yang di mock)

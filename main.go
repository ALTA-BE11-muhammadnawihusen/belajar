package main

import "fmt"

func main() {

	fmt.Printf("%q\n", `" name \ height "`)
	// "\" name \\ height \""
}

// mockery --dir=featur/user(lokasi interface yang akan di mock) --name=Datainterface(nama interface yang di mocking) --filename=UserData.go(nama file hasil generate) --structname=UserData(nama hasil strucr yang di mock)

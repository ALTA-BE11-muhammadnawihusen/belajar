package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	text := "Muhammad Nawi Husen"
	fmt.Println(len(text))

	// process encode string akan di cast ke byte terlebih dahulu di sini
	test11 := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println(test11)

	//prosess decode akan menghasilka []byte karena itu perlu di cast secara manual terlebih dahulu
	test12, _ := base64.StdEncoding.DecodeString(test11)
	test13 := string(test12)
	fmt.Println(test13)

	fmt.Println("===============================")
	//cara di bawah lwbih sulit karena terlalu banyak proses
	test21 := make([]byte, base64.StdEncoding.EncodedLen(len(text))) // akan menghasilkan wadah test21 dengan panjang yang tidak diketahui
	fmt.Println(test21)
	base64.StdEncoding.Encode(test21, []byte(text)) // mengisi wadah dengan source dari suatu sumber
	fmt.Println(test21)

	test22 := make([]byte, base64.StdEncoding.DecodedLen(len(test21))) // akan menghasilkan wadah test22
	fmt.Println(test22)
	base64.StdEncoding.Decode(test22, test21) // mengisi wadah dengan hasil decode([]byte)
	test23 := string(test22)
	fmt.Println(test23)

	fmt.Println("==============================")
	// tidak tau juga untuk apa fungsinya tapi ketika melakukan encode pada URL
	// maka lebih baik gunakan URLEncoding

	var data = "https://kalipare.com/"

	var encodedString = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println(decodedString)

}

//

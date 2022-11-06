package main

import "fmt"

func main() {
	test := "Muhammad"
	tes := []byte(test)
	fmt.Println(tes)

	te := string(tes)
	fmt.Println(te)

	test1 := int('h')
	fmt.Println(test1)
	hm := string(test1)
	fmt.Println(hm)
}

// mockery --dir=featur/user(lokasi interface yang akan di mock) --name=Datainterface(nama interface yang di mocking) --filename=UserData.go(nama file hasil generate) --structname=UserData(nama hasil strucr yang di mock)
// kita bisa melakukan langsung casting char dengan misalnya int('h') yang mana mengembalikan 104

// code fi bawah digunakan untuk assertion jika kita tidak mengetahui interface apa yang akan digunakan
// for _, val := range data {
//     switch val.(type) {
//     case string:
//         fmt.Println(val.(string))
//     case int:
//         fmt.Println(val.(int))
//     case float64:
//         fmt.Println(val.(float64))
//     case bool:
//         fmt.Println(val.(bool))
//     case []string:
//         fmt.Println(val.([]string))
//     default:
//         fmt.Println(val.(int))
//     }
// }

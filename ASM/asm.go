package main

import "fmt"

func main() {
	//salah satu cara membuat array
	var fruits = make([]string, 2)
	fruits[0] = "apple"
	fruits[1] = "manggo"

	fmt.Println(fruits) // [apple manggo]

	//akan error karena map bernilai nil sehingga tidak akan dapat diisi ketika memasukkan suatu nilai
	// var data map[string]int
	// data["one"] = 1

	// dengan menambahkan {} seperti di bawah maka map akan di buat meskipun tanpa value
	data := map[string]int{}
	data["one"] = 1

	//cara horizontal
	// var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// // cara vertical
	// var chicken2 = map[string]int{
	// 	"januari":  50,
	// 	"februari": 40,
	// }

	//menghapus suatu item yang memiliki key tertentu pada sebuah map
	// delete(chicken, "januari")

	// untuk mencari tahu kerberadaan suatu item dengan memeriksa apakah key yang di gunakan
	// ada pada sebuah map yang dicari
	// var chicken = map[string]int{"januari": 50, "februari": 40}
	// var value, isExist = chicken["mei"]
	// if isExist {
	//     fmt.Println(value)
	// } else {
	//     fmt.Println("item is not exists")
	// }

	// kita juga bisa membuat slice of map
	// var chickens = []map[string]string{
	//     map[string]string{"name": "chicken blue",   "gender": "male"},
	//     map[string]string{"name": "chicken red",    "gender": "male"},
	//     map[string]string{"name": "chicken yellow", "gender": "female"},
	// }
}

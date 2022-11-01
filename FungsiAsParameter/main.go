package main

import "fmt"

//kita juga bisa membuat sebuah alias
type namaalias func(string) bool

func Test(data []string, getdata namaalias) []string {
	return nil
}

//menggunakan fungsi sebagai parameter yang tidak pasti
func GetPP(data []string, getdata func(string) bool) []string {
	result := []string{}
	for _, item := range data {
		if temp := getdata(item); temp {
			result = append(result, item)
		}
	}

	return result
}

func main() {

	// var data = []int{1, 2, 3, 4, 5, 6, 33, 4, 43, 32, 44, 66, 11, 99, 66, 45, 15, 69, 90, 51}
	var data = []string{"Muhammad", "Nawi", "Husen", "Jack", "Phantom", "Reaper"}
	kelipatan3 := GetPP(data, func(bebas string) bool {
		return len(bebas) < 5
	})

	fmt.Println(kelipatan3)

}

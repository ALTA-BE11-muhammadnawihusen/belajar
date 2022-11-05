package main

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// fungsi ini akan menggenerati string secara acak berdasarkan huruf yang di sediakan
func randomstring(lens int) string {
	b := make([]rune, lens)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	fmt.Println(b)
	return string(b)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(1000))
	fmt.Println(rand.Intn(10000))

	test := randomstring(20)
	fmt.Println(test)

	te := []rune{233, 242, 34, 122, 244, 445, 222, 97, 225}
	fmt.Println(string(te))

}

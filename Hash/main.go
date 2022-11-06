package main

import (
	"crypto/sha1"
	"fmt"
)

//hashing merupakan proses enskripsi satu arah

func salt(text, salt string) string {
	sal := fmt.Sprintf("%s%s", text, salt)
	fmt.Println("ini yang di hash:", sal)
	sha := sha1.New()
	sha.Write([]byte(sal))
	enkr := sha.Sum(nil)
	res := fmt.Sprintf("%x", enkr)

	return res
}

func main() {
	secret := "secret"
	bumbu := "salt"

	sha := sha1.New()
	sha.Write([]byte(secret))
	fmt.Println(sha)
	enkrip := sha.Sum(nil)
	fmt.Println(enkrip) //hasil hash berupa []byte
	// fmt.Println(string(enkrip)) tapi hasil dari hash tidak bisa di stringkan kecuali dengan cara

	strhash := fmt.Sprintf("%x", enkrip)
	fmt.Println("ini adalah hasil hashing nya")
	fmt.Println(strhash)
	fmt.Println("===============================")

	test1 := salt("Kata sandi", bumbu)
	test2 := salt("Password", bumbu)
	test3 := salt("Napa ada", bumbu)

	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	if test3 == salt("Napa ada", bumbu) {
		fmt.Println("hasil sama")
	}

}
